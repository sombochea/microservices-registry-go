package api

import (
	"cubetiq/registry/controller"
	"cubetiq/registry/utils/propertymanager"
	"net"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeAPIs() error {
	router := GetRoutes()
	router.Use(cors.Default())
	return router.Run(net.JoinHostPort("", propertymanager.GetStringProperty("registry.port")))
}

func GetRoutes() *gin.Engine  {
	r := gin.Default()

	r.POST("/cubetiq/services/register/:serviceName", controller.RegisterHandler)
	r.DELETE("/cubetiq/services/unregister/:serviceName", controller.DeRegisterHandler)
	r.GET("/cubetiq/services/:serviceName", controller.QueryHandler)
	r.GET("/cubetiq/services", controller.QueryAllHandler)
	r.PUT("/cubetiq/services/:serviceName/metadata", controller.UpdateHandler)
	r.GET("/health", controller.HealthHandler)

	//KV handlers
	r.GET("/cubetiq/kv/:keyName", controller.GetKeyHandler)
	r.GET("/cubetiq/kv", controller.GetAllKeysHandler)
	r.PUT("/cubetiq/kv", controller.SetKeysHandler)
	return r
}
