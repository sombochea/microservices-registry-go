package cmd

import (
	"cubetiq/registry/api"
	"cubetiq/registry/service"
	"cubetiq/registry/utils/propertymanager"
	"cubetiq/registry/utils/redis"
	"log"
)

func Start() error {
	log.Println("Initialize configurations...")
	if err := propertymanager.InitializeConfig(); err != nil {
		return err
	}

	log.Println("Initialize redis client...")
	redis.InitializeRedis()

	log.Println("Start async health check monitoring...")
	go service.HealthCheckHandler()

	log.Println("Start the registry web server...")
	return api.InitializeAPIs()
}
