package interfaces

import (
	"context"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, request models.User) (interface{}, error)
}

type IUserHandler interface {
	RegisterUser(c *gin.Context)
}
