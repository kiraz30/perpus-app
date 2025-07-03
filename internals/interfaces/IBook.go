package interfaces

import (
	"context"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type IBookRepository interface {
	InsertNewBook(ctx context.Context, book *models.Book) error
}

type IBookService interface {
	InsertBookData(ctx context.Context, request models.Book) (interface{}, error)
}

type IBookHandler interface {
	CreateBookData(c *gin.Context)
}
