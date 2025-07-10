package interfaces

import (
	"context"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type ILendRepository interface {
	CreateNewLend(ctx context.Context, lend *models.Lend) error
}

type ILendService interface {
	CreateLendBook(ctx context.Context, request *models.Lend) (interface{}, error)
}

type ILendHandler interface {
	CreateLendBook(c *gin.Context)
}
