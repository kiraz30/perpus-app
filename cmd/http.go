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
	userV1.POST("/register", dependency.UserAPI.RegisterUser)
	userV1.POST("/login", dependency.UserAPI.Login)
	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

type Dependency struct {
	HealtCheckApi interfaces.IHealthCheckHandler
	UserAPI       interfaces.IUserHandler
}

func dependencyInject() Dependency {
	healthCheckSVC := &service.HealthCheck{}
	healthCheckApi := &api.HealthCheck{
		HealtyCheckService: healthCheckSVC,
	}

	userRepository := &repositories.UserRepository{
		DB: helpers.DB,
	}
	userSVC := &service.UserService{
		UserRepository: userRepository,
	}
	userApi := &api.UserHandler{
		UserService: userSVC,
	}

	return Dependency{
		HealtCheckApi: healthCheckApi,
		UserAPI:       userApi,
	}
}
