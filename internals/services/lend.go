package services

import (
	"context"
	"perpus-app/constants"
	"perpus-app/internals/interfaces"
	"perpus-app/internals/models"
	"time"

	"github.com/pkg/errors"
)

type LendService struct {
	LendRepository interfaces.ILendRepository
	BookRepository interfaces.IBookRepository
}

func (s *LendService) CreateLendBook(ctx context.Context, request *models.Lend) (interface{}, error) {
	var (
		statusBook int
	)
	bookCode := request.BookCode
	bookData, err := s.BookRepository.GetByBookCode(ctx, bookCode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get data book")
	}

	if bookData.BookStatus == 1 {
		return nil, errors.Wrap(err, "book not avaible")
	}
	now := time.Now()
	request.BookName = bookData.BookName
	request.DateLandBook = now
	request.DateReturnBook = now.AddDate(0, 0, 7)

	err = s.LendRepository.CreateNewLend(ctx, request)
	if err != nil {
		return nil, err
	}
	statusBook = constants.BookLend
	err = s.BookRepository.UpdateStatusBook(ctx, bookCode, statusBook)
	if err != nil {
		return nil, err
	}

	return request, nil
}
