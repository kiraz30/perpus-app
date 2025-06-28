package api

import (
	"perpus-app/helpers"
	"perpus-app/internals/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	HealtyCheckService interfaces.IHealthCheckService
}

func (api *HealthCheck) HealthCheckHandle(c *gin.Context) {
	result, err := api.HealtyCheckService.HealthCheckService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, result, nil)
}
