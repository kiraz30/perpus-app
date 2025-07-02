package repositories

import (
	"context"
	"errors"
	"perpus-app/internals/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(session).Error
}

func (r *UserRepository) GetUserSession(ctx context.Context, token string) error {

	return r.DB.Exec("SELECT * FROM user_sessions WHERE token = ?", token).Error
}

func (r *UserRepository) GetUserSessionToken(ctx context.Context, token string) (models.UserSession, error) {
	var (
		dataUserSession models.UserSession
		err             error
	)

	err = r.DB.Where("token = ?", token).First(&dataUserSession).Error
	if err != nil {
		return dataUserSession, err
	}
	if dataUserSession.ID == 0 {
		return dataUserSession, errors.New("user session not found")
	}

	return dataUserSession, nil
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}
