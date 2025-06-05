package _interface

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
    vppinterface "vpp-restapi/binapi/interface"
    vppinterface_types "vpp-restapi/binapi/interface_types"
    vppbond "vpp-restapi/binapi/bond"
    vppapi "vpp-restapi/internal/api"
)

// RegisterRoutes sets up the interface-related HTTP routes.
func RegisterRoutes(r *gin.Engine, vppClient *vppapi.VPPClient) {
    interfaceGroup := r.Group("/vpp/interfaces")
    {
        interfaceGroup.GET("", listInterfacesHandler(vppClient))
        interfaceGroup.POST("/loopback", createLoopbackHandler(vppClient))
        interfaceGroup.DELETE("/:sw_if_index", deleteInterfaceHandler(vppClient))
        // Tambahan enable/disable
        interfaceGroup.POST("/:sw_if_index/enable", enableInterfaceHandler(vppClient))
        interfaceGroup.POST("/:sw_if_index/disable", disableInterfaceHandler(vppClient))
    }
}

// Handler untuk enable interface
func enableInterfaceHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        swIfIndex, err := strconv.Atoi(c.Param("sw_if_index"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sw_if_index", "details": err.Error()})
            return
        }
        ch, err := vppClient.NewAPIChannel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer ch.Close()

        req := &vppinterface.SwInterfaceSetFlags{
            SwIfIndex: vppinterface_types.InterfaceIndex(swIfIndex),
            Flags:     vppinterface_types.IF_STATUS_API_FLAG_ADMIN_UP,
        }
        reply := &vppinterface.SwInterfaceSetFlagsReply{}
        if err := ch.SendRequest(req).ReceiveReply(reply); err != nil || reply.Retval != 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enable interface", "details": err})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Interface enabled"})
    }
}

// Handler untuk disable interface
func disableInterfaceHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        swIfIndex, err := strconv.Atoi(c.Param("sw_if_index"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sw_if_index", "details": err.Error()})
            return
        }
        ch, err := vppClient.NewAPIChannel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer ch.Close()

        req := &vppinterface.SwInterfaceSetFlags{
            SwIfIndex: vppinterface_types.InterfaceIndex(swIfIndex),
            Flags:     0, // Admin down
        }
        reply := &vppinterface.SwInterfaceSetFlagsReply{}
        if err := ch.SendRequest(req).ReceiveReply(reply); err != nil || reply.Retval != 0 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to disable interface", "details": err})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Interface disabled"})
    }
}
// listInterfacesHandler returns a handler to list all VPP interfaces.
func listInterfacesHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        ch, err := vppClient.NewAPIChannel()
        if err != nil {
            log.Printf("Failed to create API channel: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "channel creation failed", "details": err.Error()})
            return
        }
        defer ch.Close()

        interfaces := []map[string]interface{}{}
        req := &vppinterface.SwInterfaceDump{}
        reply := &vppinterface.SwInterfaceDetails{}
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
            interfaceType := inferInterfaceType(reply)
            interfaces = append(interfaces, map[string]interface{}{
                "index":     uint32(reply.SwIfIndex),
                "name":      string(reply.InterfaceName[:]),
                "type":      interfaceType,
                "admin_up":  (reply.Flags & vppinterface_types.IF_STATUS_API_FLAG_ADMIN_UP) != 0,
                "link_up":   (reply.Flags & vppinterface_types.IF_STATUS_API_FLAG_LINK_UP) != 0,
                "mtu":       reply.Mtu[0], // Use L3 MTU
            })
        }

        c.JSON(http.StatusOK, gin.H{"interfaces": interfaces})
    }
}

// createLoopbackHandler returns a handler to create a loopback interface.
func createLoopbackHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        var config struct {
            MacAddress string `json:"mac_address"`
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

        var mac [6]byte
        if config.MacAddress != "" {
            if err := parseMacAddress(config.MacAddress, &mac); err != nil {
                log.Printf("Invalid MAC address: %v", err)
                c.JSON(http.StatusBadRequest, gin.H{"error": "invalid MAC address", "details": err.Error()})
                return
            }
        }

        req := &vppinterface.CreateLoopback{
            MacAddress: mac,
        }
        reply := &vppinterface.CreateLoopbackReply{}
        if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
            log.Printf("API request failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
            return
        }
        if reply.Retval != 0 {
            log.Printf("Loopback creation failed with retval: %d", reply.Retval)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "loopback creation failed", "details": fmt.Sprintf("VPP returned non-zero retval: %d", reply.Retval)})
            return
        }

        log.Printf("Created loopback with SwIfIndex=%d", reply.SwIfIndex)
        c.JSON(http.StatusCreated, gin.H{"message": "Loopback created", "sw_if_index": reply.SwIfIndex})
    }
}

// deleteInterfaceHandler returns a handler to delete a loopback or bond interface.
func deleteInterfaceHandler(vppClient *vppapi.VPPClient) gin.HandlerFunc {
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

        // Check interface type
        reqDump := &vppinterface.SwInterfaceDump{SwIfIndex: vppinterface_types.InterfaceIndex(swIfIndex)}
        replyDump := &vppinterface.SwInterfaceDetails{}
        reqCtx := ch.SendMultiRequest(reqDump)
        isBond := false
        for {
            stop, err := reqCtx.ReceiveReply(replyDump)
            if err != nil {
                log.Printf("Dump request failed: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "dump request failed", "details": err.Error()})
                return
            }
            if stop {
                break
            }
            if strings.HasPrefix(string(replyDump.InterfaceName[:]), "BondEthernet") {
                isBond = true
            }
        }

        if isBond {
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
        } else {
            req := &vppinterface.DeleteLoopback{
                SwIfIndex: vppinterface_types.InterfaceIndex(swIfIndex),
            }
            reply := &vppinterface.DeleteLoopbackReply{}
            if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
                log.Printf("API request failed: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
                return
            }
            if reply.Retval != 0 {
                log.Printf("Loopback deletion failed with retval: %d", reply.Retval)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "loopback deletion failed", "details": fmt.Sprintf("VPP returned non-zero retval: %d", reply.Retval)})
                return
            }
            log.Printf("Deleted loopback interface %d", swIfIndex)
            c.JSON(http.StatusOK, gin.H{"message": "Loopback deleted"})
        }
    }
}
// inferInterfaceType infers the interface type based on SwInterfaceDetails.
func inferInterfaceType(details *vppinterface.SwInterfaceDetails) string {
    name := string(details.InterfaceName[:])
    switch {
    case name == "local0":
        return "local"
    case details.LinkDuplex != 0:
        return "ethernet"
    default:
        return "unknown"
    }
}

// parseMacAddress parses a MAC address string into a [6]byte array.
func parseMacAddress(macStr string, mac *[6]byte) error {
    parts := strings.Split(macStr, ":")
    if len(parts) != 6 {
        return fmt.Errorf("invalid MAC address format: %s", macStr)
    }
    for i, part := range parts {
        val, err := strconv.ParseUint(part, 16, 8)
        if err != nil {
            return fmt.Errorf("invalid MAC address part %s: %v", part, err)
        }
        mac[i] = byte(val)
    }
    return nil
}
