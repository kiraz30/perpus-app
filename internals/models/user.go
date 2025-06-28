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
