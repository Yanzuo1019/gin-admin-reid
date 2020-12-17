package router

import (
	"gin-admin-reid/api"
	"gin-admin-reid/middleware"
	"gin-admin-reid/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CorsHandler())

	r.POST("/Login", api.Login)
	r.GET("/AuthCheck", api.AuthCheck)
	r.GET("/DeviceStatus", api.DeviceStatus)
	r.POST("/Upload", api.Upload)

	network, _ := utils.Config["network"].(map[string]string)
	addr := network["addr"]
	port := network["port"]
	r.Run(addr + ":" + port)
}