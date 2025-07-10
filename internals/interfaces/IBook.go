package interfaces

import (
	"context"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type IBookRepository interface {
	InsertNewBook(ctx context.Context, book *models.Book) error
	GetListBook(ctx context.Context) ([]models.Book, error)
	GetByBookCode(ctx context.Context, BookCode string) (models.Book, error)
	UpdateStatusBook(ctx context.Context, BookCode string, status int) error
}

type IBookService interface {
	InsertBookData(ctx context.Context, request models.Book) (interface{}, error)
	GetListBook(ctx context.Context) ([]models.Book, error)
	GetByBookCode(ctx context.Context, bookCode string) (models.Book, error)
}

type IBookHandler interface {
	CreateBookData(c *gin.Context)
	GetListBook(c *gin.Context)
	GetByBookCode(c *gin.Context)
}
