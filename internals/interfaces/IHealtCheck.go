package interfaces

import "github.com/gin-gonic/gin"

type IHealthCheckHandler interface {
	HealthCheckHandle(c *gin.Context)
}

type IHealthCheckService interface {
	HealthCheckService() (string, error)
}

type IHealthCheckRepository interface {
}
