package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "vpp-restapi/internal/api"
    "vpp-restapi/internal/handler/bond"
    _interface "vpp-restapi/internal/handler/_interface"
    "vpp-restapi/internal/handler/version"
    "vpp-restapi/internal/handler/lcpng"
    "vpp-restapi/internal/handler/vlan"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    _ "vpp-restapi/docs" // <- tambahkan ini setelah generate swagger docs
)

func main() {
    log.Println("Starting VPP REST API server...")
    vppClient, err := api.NewVPPClient("/run/vpp/api.sock")
    if err != nil {
        log.Fatalf("Failed to connect to VPP: %v", err)
    }
    defer vppClient.Close()
    r := gin.Default()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    version.RegisterRoutes(r, vppClient)
    _interface.RegisterRoutes(r, vppClient)
    _interface.RegisterConfigRoutes(r, vppClient)
    bond.RegisterRoutes(r, vppClient)
    lcpng.RegisterRoutes(r, vppClient)
    vlan.RegisterRoutes(r, vppClient)
    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
