package cmd

import (
	"log"
	"perpus-app/helpers"
	"perpus-app/internals/api"
	"perpus-app/internals/interfaces"
	"perpus-app/internals/repositories"
	service "perpus-app/internals/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()
	r.GET("/healty", dependency.HealtCheckApi.HealthCheckHandle)

	userV1 := r.Group("/v1/user")
	userV1.POST("/register", dependency.RegisterAPI.RegisterUser)
	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

type Dependency struct {
	HealtCheckApi interfaces.IHealthCheckHandler
	RegisterAPI   interfaces.IUserHandler
}

func dependencyInject() Dependency {
	healthCheckSVC := &service.HealthCheck{}
	healthCheckApi := &api.HealthCheck{
		HealtyCheckService: healthCheckSVC,
	}

	userRepository := &repositories.UserRepository{
		DB: helpers.DB,
	}
	registerSVC := &service.UserService{
		UserRepository: userRepository,
	}
	registerApi := &api.UserHandler{
		UserService: registerSVC,
	}

	return Dependency{
		HealtCheckApi: healthCheckApi,
		RegisterAPI:   registerApi,
	}
}
