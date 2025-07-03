package interfaces

import (
	"context"
	"perpus-app/helpers"
	"perpus-app/internals/models"

	"github.com/gin-gonic/gin"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	GetUserSessionToken(ctx context.Context, token string) (models.UserSession, error)
	DeleteUserSession(ctx context.Context, token string) error
	GetUserRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
	UpdateRefreshToken(ctx context.Context, token, refreshToken string) error
}

type IUserService interface {
	RegisterUser(ctx context.Context, request models.User) (interface{}, error)
	Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}

type IUserHandler interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
}
