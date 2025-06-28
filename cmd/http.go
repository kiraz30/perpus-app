package cmd

import (
	"log"
	"perpus-app/helpers"
	"perpus-app/internals/api"
	"perpus-app/internals/interfaces"
	service "perpus-app/internals/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()

	r.GET("/healty", dependency.HealtCheckApi.HealthCheckHandle)
	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

type Dependency struct {
	HealtCheckApi interfaces.IHealthCheckHandler
}

func dependencyInject() Dependency {
	healthCheckSVC := &service.HealthCheck{}
	healthCheckApi := &api.HealthCheck{
		HealtyCheckService: healthCheckSVC,
	}

	return Dependency{
		HealtCheckApi: healthCheckApi,
	}
}
