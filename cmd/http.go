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

	userV1WithAuth := userV1.Use()
	userV1WithAuth.POST("/logout", dependency.MiddlewareValidateAuth, dependency.UserAPI.Logout)
	userV1WithAuth.PUT("/refresh-token", dependency.MiddlewareValidateAuth, dependency.UserAPI.RefreshToken)

	bookV1 := r.Group("/v1/book")
	bookV1WithAuth := bookV1.Use()
	bookV1WithAuth.POST("/create", dependency.MiddlewareValidateAuth, dependency.BookAPI.CreateBookData)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

type Dependency struct {
	HealtCheckApi  interfaces.IHealthCheckHandler
	UserAPI        interfaces.IUserHandler
	UserRepository interfaces.IUserRepository
	BookAPI        interfaces.IBookHandler
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

	bookRepository := &repositories.BookRepository{
		DB: helpers.DB,
	}
	bookSVC := &service.BookService{
		BookRepository: bookRepository,
	}
	bookAPI := &api.BookApi{
		BookService: bookSVC,
	}

	return Dependency{
		HealtCheckApi:  healthCheckApi,
		UserAPI:        userApi,
		UserRepository: userRepository,
		BookAPI:        bookAPI,
	}
}
