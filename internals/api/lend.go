package api

import (
	"net/http"
	"perpus-app/constants"
	"perpus-app/helpers"
	"perpus-app/internals/interfaces"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type LendApi struct {
	LendService interfaces.ILendService
}

func (api *LendApi) CreateLendBook(c *gin.Context) {
	var (
		log     = helpers.Logger
		request models.Lend
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Error("failed to get request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	err = request.Validate()
	if err != nil {
		log.Error("failed to validate", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	response, err := api.LendService.CreateLendBook(c.Request.Context(), &request)
	if err != nil {
		log.Error("failed to create lend book", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedInternalServer, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccesMessage, response)

}
