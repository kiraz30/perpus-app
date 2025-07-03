package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Book struct {
	ID        int       `json:"id"`
	BookName  string    `json:"book_name" gorm:"column:book_name;type:varchar(255)" validate:"required"`
	Genre     string    `json:"genre" gorm:"column:genre;type:varchar(100)" validate:"required"`
	Author    string    `json:"author" gorm:"column:author;type:varchar(100)" validate:"required"`
	Publisher string    `json:"publisher" gorm:"column:publisher;type:varchar(100)" validate:"required"`
	BookCode  string    `json:"book_code" gorm:"column:book_code;type:varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*Book) TableName() string {
	return "books"
}

func (l Book) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
