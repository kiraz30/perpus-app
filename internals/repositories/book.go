package repositories

import (
	"context"
	"perpus-app/internals/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) InsertNewBook(ctx context.Context, book *models.Book) error {
	return r.DB.Create(book).Error
}
