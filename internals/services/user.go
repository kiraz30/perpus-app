package services

import (
	"context"
	"perpus-app/internals/interfaces"
	"perpus-app/internals/models"

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
