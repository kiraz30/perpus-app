package services

import (
	"context"
	"perpus-app/helpers"
	"perpus-app/internals/interfaces"
	"perpus-app/internals/models"
	"time"

	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository interfaces.IUserRepository
}

func (s *UserService) RegisterUser(ctx context.Context, request models.User) (interface{}, error) {

	//hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	request.Password = string(hashPassword)

	err = s.UserRepository.InsertNewUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	response := request
	response.Password = ""
	return response, nil

}

func (s *UserService) Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error) {
	var (
		response models.LoginResponse
		now      = time.Now()
		// tokenType = "token"
	)
	userData, err := s.UserRepository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return response, errors.Wrap(err, "failed to get user data")
	}
	if userData.ID == 0 {
		return response, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(request.Password))
	if err != nil {
		return response, errors.Wrap(err, "password does not match")
	}

	token, err := helpers.GenerateToken(ctx, userData.ID, userData.Username, userData.Email, "token", now)
	if err != nil {
		return response, errors.Wrap(err, "failed to generate token")
	}

	userSession := &models.UserSession{
		UserID:                userData.ID,
		Token:                 token,
		RefreshToken:          "",
		TokenExpiresAt:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpiresAt: now.Add(helpers.MapTypeToken["refresh_token"]),
	}

	err = s.UserRepository.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return response, errors.Wrap(err, "failed to insert user session")

	}
	response.UserID = userData.ID
	response.Username = userData.Username
	response.Email = userData.Email
	response.Token = token

	return response, nil
}
