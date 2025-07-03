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
