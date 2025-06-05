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
func RegisterRoutes(r *gin.Engine, vppClient *vppapi.VPPClient) {
    bondGroup := r.Group("/vpp/bonds")
    {
        bondGroup.GET("", listBondsHandler(vppClient))
        bondGroup.POST("", createBondHandler(vppClient))
        bondGroup.POST("/:sw_if_index/member", addBondMemberHandler(vppClient))
        bondGroup.DELETE("/:sw_if_index", deleteBondHandler(vppClient))
    }
}

// listBondsHandler returns a handler to list all bond interfaces.
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
                "active_slaves": reply.ActiveSlaves,
                "slaves":    reply.Slaves,
            })
        }

        c.JSON(http.StatusOK, gin.H{"bonds": bonds})
    }
}

// createBondHandler returns a handler to create a bond interface.
func createBondHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        var config struct {
            Mode       string   `json:"mode"`
            LbMode     string   `json:"lb_mode"`
            ID         *uint32  `json:"id,omitempty"`
            Interfaces []uint32 `json:"interfaces"`
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
        default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mode", "details": "supported modes: lacp, round-robin"})
            return
        }

        var lb vppbond.BondLbAlgo
        switch config.LbMode {
        case "l2":
            lb = vppbond.BOND_API_LB_ALGO_L2
        case "l23":
            lb = vppbond.BOND_API_LB_ALGO_L23
        case "l34":
            lb = vppbond.BOND_API_LB_ALGO_L34
        case "":
            lb = vppbond.BOND_API_LB_ALGO_L2 // default jika kosong
        default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lb_mode", "details": "supported lb_mode: l2, l23, l34"})
            return
        }

        req := &vppbond.BondCreate{
            Mode: mode,
            Lb:   lb,
        }
        // Jika ID diberikan, gunakan, jika tidak biarkan default (0xFFFFFFFF) supaya VPP auto assign
        if config.ID != nil {
            req.ID = *config.ID
        } else {
            req.ID = 0xFFFFFFFF // default (biar VPP auto)
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
// addBondMemberHandler returns a handler to add a member to a bond interface.
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

// deleteBondHandler returns a handler to delete a bond interface.
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
