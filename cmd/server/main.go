package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "vpp-restapi/internal/api"
    "vpp-restapi/internal/handler/bond"
    _interface "vpp-restapi/internal/handler/_interface"
    "vpp-restapi/internal/handler/version"
)

func main() {
    log.Println("Starting VPP REST API server...")
    vppClient, err := api.NewVPPClient("/run/vpp/api.sock")
    if err != nil {
        log.Fatalf("Failed to connect to VPP: %v", err)
    }
    defer vppClient.Close()
    r := gin.Default()
    version.RegisterRoutes(r, vppClient)
    _interface.RegisterRoutes(r, vppClient)
    bond.RegisterRoutes(r, vppClient)
    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
