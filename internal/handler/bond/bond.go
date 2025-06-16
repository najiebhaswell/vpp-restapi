package bond

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    vppbond "vpp-restapi/binapi/bond"
    vppinterface_types "vpp-restapi/binapi/interface_types"
    vppapi "vpp-restapi/internal/api"
)

// RegisterRoutes sets up the bond-related HTTP routes.
// Sekarang menerima gin.IRoutes, tanpa membuat group di sini.
func RegisterRoutes(r gin.IRoutes, vppClient *vppapi.VPPClient) {
    r.GET("/bonds", listBondsHandler(vppClient))
    r.POST("/bonds", createBondHandler(vppClient))
    r.POST("/bonds/:sw_if_index/member", addBondMemberHandler(vppClient))
    r.DELETE("/bonds/:sw_if_index", deleteBondHandler(vppClient))
}

// @Summary List Bond Interfaces
// @Description List all bond interfaces.
// @Tags bonds
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /vpp/bonds [get]
func listBondsHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        ch, err := vppClient.NewAPIChannel()
        if err != nil {
            log.Printf("Failed to create API channel: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "channel creation failed", "details": err.Error()})
            return
        }
        defer ch.Close()

        bonds := []map[string]interface{}{}
        req := &vppbond.SwInterfaceBondDump{}
        reply := &vppbond.SwInterfaceBondDetails{}
        reqCtx := ch.SendMultiRequest(req)
        for {
            stop, err := reqCtx.ReceiveReply(reply)
            if err != nil {
                log.Printf("API request failed: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
                return
            }
            if stop {
                break
            }
            bonds = append(bonds, map[string]interface{}{
                "index":     uint32(reply.SwIfIndex),
                "name":      reply.InterfaceName,
                "mode":      reply.Mode.String(),
                "lb_algo":   reply.Lb.String(),
                "admin_up":  nil, // Not available in SwInterfaceBondDetails
                "link_up":   nil, // Not available in SwInterfaceBondDetails
                "members":   reply.Slaves,
                "active":    reply.ActiveSlaves,
            })
        }

        c.JSON(http.StatusOK, gin.H{"bonds": bonds})
    }
}

// @Summary Create Bond Interface
// @Description Create a new bond interface (multi-bond supported, id required for unique bond).
// @Tags bonds
// @Accept json
// @Produce json
// @Param body body object true "Bond Config {mode: string, interfaces: []int, id: int (optional), mac_address: string (optional), lb: string (optional), numa_only: bool (optional)}"
// @Success 201 {object} map[string]interface{}
// @Failure 400,500 {object} map[string]interface{}
// @Router /vpp/bonds [post]
func createBondHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        var config struct {
            Mode        string   `json:"mode"`
            Interfaces  []uint32 `json:"interfaces"`
            Id          *uint32  `json:"id,omitempty"`
            MacAddress  string   `json:"mac_address,omitempty"`
            Lb          string   `json:"lb,omitempty"`
            NumaOnly    *bool    `json:"numa_only,omitempty"`
        }
        if err := c.ShouldBindJSON(&config); err != nil {
            log.Printf("Invalid request body: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
            return
        }

        ch, err := vppClient.NewAPIChannel()
        if err != nil {
            log.Printf("Failed to create API channel: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "channel creation failed", "details": err.Error()})
            return
        }
        defer ch.Close()

        var mode vppbond.BondMode
        switch config.Mode {
        case "lacp":
            mode = vppbond.BOND_API_MODE_LACP
        case "round-robin":
            mode = vppbond.BOND_API_MODE_ROUND_ROBIN
        case "xor":
            mode = vppbond.BOND_API_MODE_XOR
        case "active-backup":
            mode = vppbond.BOND_API_MODE_ACTIVE_BACKUP
        case "broadcast":
            mode = vppbond.BOND_API_MODE_BROADCAST
        default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mode", "details": "supported modes: lacp, round-robin, xor, active-backup, broadcast"})
            return
        }

        req := &vppbond.BondCreate{
            Mode: mode,
        }
        if config.Id != nil {
            req.ID = *config.Id
        }
        if config.MacAddress != "" {
            copy(req.MacAddress[:], parseMacString(config.MacAddress))
            req.UseCustomMac = true
        }
        if config.Lb != "" {
            switch config.Lb {
            case "l2":
                req.Lb = vppbond.BOND_API_LB_ALGO_L2
            case "l34":
                req.Lb = vppbond.BOND_API_LB_ALGO_L34
            case "l23":
                req.Lb = vppbond.BOND_API_LB_ALGO_L23
            case "rr":
                req.Lb = vppbond.BOND_API_LB_ALGO_RR
            case "bc":
                req.Lb = vppbond.BOND_API_LB_ALGO_BC
            case "ab":
                req.Lb = vppbond.BOND_API_LB_ALGO_AB
            default:
                // if unknown, ignore
            }
        }
        if config.NumaOnly != nil {
            req.NumaOnly = *config.NumaOnly
        }

        reply := &vppbond.BondCreateReply{}
        if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
            log.Printf("API request failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
            return
        }
        if reply.Retval != 0 {
            log.Printf("Bond creation failed with retval: %d", reply.Retval)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "bond creation failed", "details": fmt.Sprintf("VPP returned non-zero retval: %d", reply.Retval)})
            return
        }

        swIfIndex := reply.SwIfIndex
        for _, memberIdx := range config.Interfaces {
            reqAdd := &vppbond.BondAddMember{
                SwIfIndex:     vppinterface_types.InterfaceIndex(memberIdx),
                BondSwIfIndex: swIfIndex,
                IsPassive:     false,
                IsLongTimeout: false,
            }
            replyAdd := &vppbond.BondAddMemberReply{}
            if err := ch.SendRequest(reqAdd).ReceiveReply(replyAdd); err != nil {
                log.Printf("Add member failed: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "add member failed", "details": err.Error()})
                return
            }
            if replyAdd.Retval != 0 {
                log.Printf("Add member failed with retval: %d", replyAdd.Retval)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "add member failed", "details": fmt.Sprintf("VPP returned non-zero retval: %d", replyAdd.Retval)})
                return
            }
        }

        log.Printf("Created bond with SwIfIndex=%d", swIfIndex)
        c.JSON(http.StatusCreated, gin.H{"message": "Bond created", "sw_if_index": swIfIndex})
    }
}

// Helper: convert MAC string to [6]byte for VPP
func parseMacString(mac string) []byte {
    var b [6]byte
    fmt.Sscanf(mac, "%x:%x:%x:%x:%x:%x", &b[0], &b[1], &b[2], &b[3], &b[4], &b[5])
    return b[:]
}

// @Summary Add Bond Member
// @Description Add a member to a bond interface.
// @Tags bonds
// @Accept json
// @Produce json
// @Param sw_if_index path int true "Bond Interface Index"
// @Param body body object true "Member {member_index: int}"
// @Success 200 {object} map[string]interface{}
// @Failure 400,500 {object} map[string]interface{}
// @Router /vpp/bonds/{sw_if_index}/member [post]
func addBondMemberHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        swIfIndex, err := strconv.Atoi(c.Param("sw_if_index"))
        if err != nil {
            log.Printf("Invalid sw_if_index: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sw_if_index", "details": err.Error()})
            return
        }

        var config struct {
            MemberIndex uint32 `json:"member_index"`
        }
        if err := c.ShouldBindJSON(&config); err != nil {
            log.Printf("Invalid request body: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
            return
        }

        ch, err := vppClient.NewAPIChannel()
        if err != nil {
            log.Printf("Failed to create API channel: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "channel creation failed", "details": err.Error()})
            return
        }
        defer ch.Close()

        req := &vppbond.BondAddMember{
            SwIfIndex:     vppinterface_types.InterfaceIndex(config.MemberIndex),
            BondSwIfIndex: vppinterface_types.InterfaceIndex(swIfIndex),
            IsPassive:     false,
            IsLongTimeout: false,
        }
        reply := &vppbond.BondAddMemberReply{}
        if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
            log.Printf("API request failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
            return
        }
        if reply.Retval != 0 {
            log.Printf("Add member failed with retval: %d", reply.Retval)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "add member failed", "details": fmt.Sprintf("VPP returned non-zero retval: %d", reply.Retval)})
            return
        }

        log.Printf("Added member %d to bond %d", config.MemberIndex, swIfIndex)
        c.JSON(http.StatusOK, gin.H{"message": "Member added"})
    }
}

// @Summary Delete Bond Interface
// @Description Delete a bond interface by sw_if_index.
// @Tags bonds
// @Param sw_if_index path int true "Bond Interface Index"
// @Success 200 {object} map[string]interface{}
// @Failure 400,500 {object} map[string]interface{}
// @Router /vpp/bonds/{sw_if_index} [delete]
func deleteBondHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        swIfIndex, err := strconv.Atoi(c.Param("sw_if_index"))
        if err != nil {
            log.Printf("Invalid sw_if_index: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sw_if_index", "details": err.Error()})
            return
        }

        ch, err := vppClient.NewAPIChannel()
        if err != nil {
            log.Printf("Failed to create API channel: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "channel creation failed", "details": err.Error()})
            return
        }
        defer ch.Close()

        req := &vppbond.BondDelete{
            SwIfIndex: vppinterface_types.InterfaceIndex(swIfIndex),
        }
        reply := &vppbond.BondDeleteReply{}
        if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
            log.Printf("API request failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
            return
        }
        if reply.Retval != 0 {
            log.Printf("Bond deletion failed with retval: %d", reply.Retval)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "bond deletion failed", "details": fmt.Sprintf("VPP returned non-zero retval: %d", reply.Retval)})
            return
        }

        log.Printf("Deleted bond interface %d", swIfIndex)
        c.JSON(http.StatusOK, gin.H{"message": "Bond deleted"})
    }
}
