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

func (r *BookRepository) GetListBook(ctx context.Context) ([]models.Book, error) {
	var response []models.Book
	err := r.DB.Debug().Order("id DESC").Find(&response).Error
	return response, err
}

func (r *BookRepository) GetByBookCode(ctx context.Context, BookCode string) (models.Book, error) {
	var response models.Book
	err := r.DB.Debug().Where("book_code = ?", BookCode).Last(&response).Error
	return response, err
}

func (r *BookRepository) UpdateStatusBook(ctx context.Context, BookCode string, status int) error {
	return r.DB.Exec("UPDATE books SET book_status = ? WHERE book_code = ?", status, BookCode).Error
}
