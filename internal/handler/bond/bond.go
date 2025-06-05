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
        req := &vppbond.BondDump{}
        reply := &vppbond.BondDetails{}
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
                "name":      string(reply.InterfaceName[:]),
                "mode":      reply.Mode.String(),
                "admin_up":  reply.IsEnabled,
                "link_up":   reply.IsUp,
                "members":   len(reply.ActiveMembers),
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

        req := &vppbond.BondCreate{
            Mode:      mode,
            UseCarrier: true,
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
            reqAdd := &vppbond.SwInterfaceBondEnslave{
                SwIfIndex:  vppinterface_types.InterfaceIndex(memberIdx),
                BondSwIfIndex: swIfIndex,
                IsPassive:  false,
                IsLongTimeout: false,
            }
            replyAdd := &vppbond.SwInterfaceBondEnslaveReply{}
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

        req := &vppbond.SwInterfaceBondEnslave{
            SwIfIndex:  vppinterface_types.InterfaceIndex(config.MemberIndex),
            BondSwIfIndex: vppinterface_types.InterfaceIndex(swIfIndex),
            IsPassive:  false,
            IsLongTimeout: false,
        }
        reply := &vppbond.SwInterfaceBondEnslaveReply{}
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
