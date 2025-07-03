package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" gorm:"column:username;type:varchar(20)" validate:"required"`
	Email     string    `json:"email" gorm:"column:email;type:varchar(100)" validate:"required,email"`
	Password  string    `json:"password,omitempty" gorm:"column:password;type:varchar(255)" validate:"required" `
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*User) TableName() string {
	return "users"
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

// User Session
type UserSession struct {
	ID                    int       `json:"id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	UserID                int       `json:"user_id" gorm:"column:user_id;type:int"`
	Token                 string    `json:"token" gorm:"column:token;type:text" validate:"required"`
	RefreshToken          string    `json:"refresh_token" gorm:"column:refresh_token;type:text" validate:"required"`
	TokenExpiresAt        time.Time `json:"-" validate:"required"`
	RefreshTokenExpiresAt time.Time `json:"-" validate:"required"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

// login
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (l LoginRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type LoginResponse struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}
