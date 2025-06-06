package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    api "go.fd.io/govpp/api"
    "go.fd.io/govpp/adapter/socketclient"
    "go.fd.io/govpp/core"
    vppbond "vpp-restapi/binapi/bond"
    vppinterface "vpp-restapi/binapi/interface"
    "vpp-restapi/binapi/interface_types"
    vpptap "vpp-restapi/binapi/tapv2"
    "vpp-restapi/binapi/vpe"
)

const vppSocket = "/run/vpp/api.sock"

func main() {
    log.Println("Starting VPP REST API server...")

    adapter := socketclient.NewVppClient(vppSocket)
    conn, err := core.Connect(adapter)
    if err != nil {
        log.Fatalf("Failed to connect to VPP: %v", err)
    }
    defer conn.Disconnect()

    r := gin.Default()

    // Endpoint untuk mendapatkan versi VPP
    r.GET("/vpp/version", versionHandler(conn))

    // Endpoint untuk mendapatkan daftar antarmuka
    r.GET("/vpp/interfaces", interfacesHandler(conn))

    // Endpoint untuk konfigurasi antarmuka (create/delete loopback, delete interface)
    r.POST("/vpp/interfaces/:index/config", configInterfaceHandler(conn))

    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Gagal menjalankan server: %v", err)
    }
}

func versionHandler(conn *core.Connection) gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Handling /vpp/version request")
        ch, err := conn.NewAPIChannel()
        if err != nil {
            log.Printf("Failed to create API channel: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "channel creation failed", "details": err.Error()})
            return
        }
        defer ch.Close()

        req := &vpe.ShowVersion{}
        reply := &vpe.ShowVersionReply{}
        if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
            log.Printf("API request failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "version":    reply.Version,
            "build_date": reply.BuildDate,
        })
    }
}

func interfacesHandler(conn *core.Connection) gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Handling /vpp/interfaces request")
        ch, err := conn.NewAPIChannel()
        if err != nil {
            log.Printf("Failed to create API channel: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "channel creation failed", "details": err.Error()})
            return
        }
        defer ch.Close()

        // Cache nama antarmuka
        nameCache := make(map[uint32]string)

        req := &vppinterface.SwInterfaceDump{}
        reply := &vppinterface.SwInterfaceDetails{}
        interfacesList := []map[string]interface{}{}

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
            // Gunakan fallback untuk nama
            swIfIndex := uint32(reply.SwIfIndex)
            name := fmt.Sprintf("interface-%d", swIfIndex)
            // Asumsikan nama berdasarkan SwIfIndex dan konteks VPP
            if swIfIndex == 0 {
                name = "local0"
            }
            log.Printf("Interface Details: SwIfIndex=%d, Name=%s, Type=%v, Flags=%d, MTU=%d", swIfIndex, name, reply.Type, reply.Flags, reply.LinkMtu)
            nameCache[swIfIndex] = name
            interfacesList = append(interfacesList, map[string]interface{}{
                "index":     swIfIndex,
                "name":      nameCache[swIfIndex],
                "admin_up":  reply.Flags&interface_types.IF_STATUS_API_FLAG_ADMIN_UP != 0,
                "link_up":   reply.Flags&interface_types.IF_STATUS_API_FLAG_LINK_UP != 0,
                "mtu":       reply.LinkMtu,
            })
        }

        c.JSON(http.StatusOK, gin.H{
            "interfaces": interfacesList,
        })
    }
}

func getInterfaceDetails(ch api.Channel, swIfIndex uint32) (*vppinterface.SwInterfaceDetails, error) {
    req := &vppinterface.SwInterfaceDump{
        SwIfIndex: interface_types.InterfaceIndex(swIfIndex),
    }
    reply := &vppinterface.SwInterfaceDetails{}
    reqCtx := ch.SendMultiRequest(req)
    stop, err := reqCtx.ReceiveReply(reply)
    if err != nil {
        return nil, fmt.Errorf("failed to dump interface: %v", err)
    }
    if stop {
        return nil, fmt.Errorf("interface %d not found", swIfIndex)
    }
    log.Printf("Interface %d: Type=%v", swIfIndex, reply.Type)
    return reply, nil
}

func inferInterfaceType(swIfIndex uint32) (string, string) {
    // Mengembalikan jenis antarmuka dan nama berdasarkan SwIfIndex dan konvensi VPP
    if swIfIndex == 0 {
        return "local", "local0"
    }
    // Asumsikan jenis berdasarkan SwIfIndex dari output vppctl
    // Sesuaikan dengan daftar antarmuka: loop0 (1), BondEthernet0 (2), tap0 (3), loop1 (4)
    switch swIfIndex {
    case 1:
        return "loopback", "loop0"
    case 2:
        return "bond", "BondEthernet0"
    case 3:
        return "tap", "tap0"
    case 4:
        return "loopback", "loop1"
    default:
        return "unknown", fmt.Sprintf("interface-%d", swIfIndex)
    }
}

func configInterfaceHandler(conn *core.Connection) gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Handling /vpp/interfaces/:index/config request")
        swIfIndex, err := strconv.Atoi(c.Param("index"))
        if err != nil {
            log.Printf("Invalid interface index: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index", "details": err.Error()})
            return
        }

        var config struct {
            Action string `json:"action"` // create_loopback, delete_interface
        }
        if err := c.ShouldBindJSON(&config); err != nil {
            log.Printf("Invalid request body: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body", "details": err.Error()})
            return
        }

        ch, err := conn.NewAPIChannel()
        if err != nil {
            log.Printf("Failed to create API channel: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "channel creation failed", "details": err.Error()})
            return
        }
        defer ch.Close()

        switch config.Action {
        case "create_loopback":
            req := &vppinterface.CreateLoopback{}
            reply := &vppinterface.CreateLoopbackReply{}
            if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
                log.Printf("API request failed: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
                return
            }
            log.Printf("Created loopback with SwIfIndex=%d", reply.SwIfIndex)
            c.JSON(http.StatusCreated, gin.H{"message": "Loopback created", "sw_if_index": reply.SwIfIndex})

        case "delete_interface":
            details, err := getInterfaceDetails(ch, uint32(swIfIndex))
            if err != nil {
                log.Printf("Failed to get interface details: %v", err)
                c.JSON(http.StatusBadRequest, gin.H{"error": "invalid interface", "details": err.Error()})
                return
            }
            // Tentukan jenis dan nama antarmuka
            ifType, name := inferInterfaceType(uint32(swIfIndex))
            // Cegah penghapusan local0
            if ifType == "local" {
                log.Printf("Cannot delete local0 interface")
                c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete", "details": "local0 interface cannot be deleted"})
                return
            }
            // Panggil API penghapusan berdasarkan tipe
            switch ifType {
            case "loopback":
                req := &vppinterface.DeleteLoopback{
                    SwIfIndex: interface_types.InterfaceIndex(swIfIndex),
                }
                reply := &vppinterface.DeleteLoopbackReply{}
                if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
                    log.Printf("API request failed: %v", err)
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
                    return
                }
                log.Printf("Deleted loopback interface %d (%s)", swIfIndex, name)
                c.JSON(http.StatusOK, gin.H{"message": "Interface deleted", "name": name})
            case "bond":
                req := &vppbond.BondDelete{
                    SwIfIndex: interface_types.InterfaceIndex(swIfIndex),
                }
                reply := &vppbond.BondDeleteReply{}
                if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
                    log.Printf("API request failed: %v", err)
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
                    return
                }
                log.Printf("Deleted bond interface %d (%s)", swIfIndex, name)
                c.JSON(http.StatusOK, gin.H{"message": "Interface deleted", "name": name})
            case "tap":
                req := &vpptap.TapDeleteV2{
                    SwIfIndex: interface_types.InterfaceIndex(swIfIndex),
                }
                reply := &vpptap.TapDeleteV2Reply{}
                if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
                    log.Printf("API request failed: %v", err)
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
                    return
                }
                log.Printf("Deleted tap interface %d (%s)", swIfIndex, name)
                c.JSON(http.StatusOK, gin.H{"message": "Interface deleted", "name": name})
            default:
                log.Printf("Unsupported interface type for deletion: %s", ifType)
                c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported interface", "details": fmt.Sprintf("interface %s (type %s) cannot be deleted", name, ifType)})
                return
            }

        default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid action", "details": "supported actions: create_loopback, delete_interface"})
        }
    }
}
