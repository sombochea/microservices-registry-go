package controller

import (
	"cubetiq/registry/model"
	"cubetiq/registry/utils/redis"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeRegisterHandler(ctx *gin.Context) {

	serviceName := ctx.Param("serviceName")
	result, err := redis.HDel("services", serviceName)
	log.Println("Deregister result:", result)
	if err != nil || result == 0{
		ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1003, "Failed to deregister service or service not found"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":"1000",
		"status_description":"Successfully removed",
	})
}