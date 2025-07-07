package models

import "time"

type Lend struct {
	ID             int       `json:"id"`
	BookName       string    `json:"book_name" gorm:"column:book_name;type:varchar(255)" validate:"required"`
	BookCode       string    `json:"book_code" gorm:"column:book_code;type:varchar(255)" validate:"required"`
	UserId         int       `json:"user_id" gorm:"column:user_id;type:int(100)" validate:"required"`
	UserFullName   string    `json:"user_full_name" gorm:"column:user_full_name;type:int(100)" validate:"required"`
	UserEmail      string    `json:"user_email" gorm:"column:user_email;type:varchar(25)" validate:"required"`
	DateLandBook   time.Time `json:"date_land_book" gorm:"column:date_land_book" validate:"required"`
	DateReturnBook time.Time `json:"date_return_book" gorm:"column:date_return_book" validate:"required"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
