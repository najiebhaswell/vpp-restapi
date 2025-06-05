package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "go.fd.io/govpp/adapter/socketclient"
    "go.fd.io/govpp/core"
    vppinterface "vpp-restapi/binapi/interface"
    "vpp-restapi/binapi/interface_types"
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

    // Endpoint untuk konfigurasi antarmuka (create/delete loopback)
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
            // Gunakan InterfaceName jika tersedia, jika tidak gunakan fallback
            swIfIndex := uint32(reply.SwIfIndex)
            name := string(reply.InterfaceName)
            if name == "" {
                name = fmt.Sprintf("interface-%d", swIfIndex)
            }
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
            Action string `json:"action"` // create_loopback, delete_loopback
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
            c.JSON(http.StatusCreated, gin.H{"message": "Loopback created", "sw_if_index": reply.SwIfIndex})

        case "delete_loopback":
            req := &vppinterface.DeleteLoopback{
                SwIfIndex: interface_types.InterfaceIndex(swIfIndex),
            }
            reply := &vppinterface.DeleteLoopbackReply{}
            if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
                log.Printf("API request failed: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed", "details": err.Error()})
                return
            }
            c.JSON(http.StatusOK, gin.H{"message": "Loopback deleted"})

        default:
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid action", "details": "supported actions: create_loopback, delete_loopback"})
        }
    }
}
