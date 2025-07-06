package api

import (
	"net/http"
	"perpus-app/constants"
	"perpus-app/helpers"
	"perpus-app/internals/interfaces"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type BookApi struct {
	BookService interfaces.IBookService
}

func (api *BookApi) CreateBookData(c *gin.Context) {
	var (
		log     = helpers.Logger
		request models.Book
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

	response, err := api.BookService.InsertBookData(c.Request.Context(), request)
	if err != nil {
		log.Error("failed to insert book", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedInternalServer, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccesMessage, response)

}

func (api *BookApi) GetListBook(c *gin.Context) {
	var log = helpers.Logger

	response, err := api.BookService.GetListBook(c.Request.Context())
	if err != nil {
		log.Errorf("Failed to get list Book ")
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccesMessage, response)
}

func (api *BookApi) GetByBookCode(c *gin.Context) {
	var log = helpers.Logger

	bookCode := c.Param("bookCode")

	response, err := api.BookService.GetByBookCode(c.Request.Context(), bookCode)
	if err != nil {
		log.Errorf("Failed to get Book ")
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccesMessage, response)
}
