package api

import (
	"net/http"
	"perpus-app/constants"
	"perpus-app/helpers"
	"perpus-app/internals/interfaces"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService interfaces.IUserService
}

func (api *UserHandler) RegisterUser(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	request := models.User{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Info("Failed to parse request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	err = request.Validate()
	if err != nil {
		log.Info("Failed to validate", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	response, err := api.UserService.RegisterUser(c.Request.Context(), request)
	if err != nil {
		log.Error("Failed to register user", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccesMessage, response)
	return
}

func (api *UserHandler) Login(c *gin.Context) {
	var (
		log      = helpers.Logger
		request  = models.LoginRequest{}
		response = models.LoginResponse{}
	)

	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		log.Info("Failed to parse request :", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	err = request.Validate()
	if err != nil {
		log.Info("Failed to validate :", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	response, err = api.UserService.Login(c.Request.Context(), request)
	if err != nil {
		log.Info("Failed to login :", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccesMessage, response)
	return
}

func (api *UserHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	token := c.Request.Header.Get("Authorization")

	err := api.UserService.Logout(c.Request.Context(), token)
	if err != nil {
		log.Info("Failed to logout :", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedInternalServer, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccesMessage, nil)
}
