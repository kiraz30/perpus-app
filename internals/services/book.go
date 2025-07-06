package services

import (
	"context"
	"perpus-app/helpers"
	"perpus-app/internals/interfaces"
	"perpus-app/internals/models"
)

type BookService struct {
	BookRepository interfaces.IBookRepository
}

func (s *BookService) InsertBookData(ctx context.Context, request models.Book) (interface{}, error) {

	request.BookCode = helpers.GenerateBookCode()
	err := s.BookRepository.InsertNewBook(ctx, &request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (s *BookService) GetListBook(ctx context.Context) ([]models.Book, error) {
	return s.BookRepository.GetListBook(ctx)
}

func (s *BookService) GetByBookCode(ctx context.Context, bookCode string) (models.Book, error) {
	return s.BookRepository.GetByBookCode(ctx, bookCode)
}
