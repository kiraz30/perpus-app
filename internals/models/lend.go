package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Lend struct {
	ID             int       `json:"id"`
	BookName       string    `json:"book_name" gorm:"column:book_name;type:varchar(255)"`
	BookCode       string    `json:"book_code" gorm:"column:book_code;type:varchar(255)" validate:"required"`
	UserId         int       `json:"user_id" gorm:"column:user_id;type:int(100)"`
	UserFullName   string    `json:"user_full_name" gorm:"column:user_full_name;type:varchar(100)"`
	UserEmail      string    `json:"user_email" gorm:"column:user_email;type:varchar(25)"`
	DateLandBook   time.Time `json:"date_land_book" gorm:"column:date_land_book"`
	DateReturnBook time.Time `json:"date_return_book" gorm:"column:date_return_book"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

func (*Lend) TableName() string {
	return "lend"
}

func (l Lend) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
