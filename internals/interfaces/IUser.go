package interfaces

import (
	"context"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, request models.User) (interface{}, error)
	Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error)
}

type IUserHandler interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
}
