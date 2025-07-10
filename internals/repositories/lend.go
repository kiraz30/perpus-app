package repositories

import (
	"context"
	"perpus-app/internals/models"

	"gorm.io/gorm"
)

type LendRepository struct {
	DB *gorm.DB
}

func (r *LendRepository) CreateNewLend(ctx context.Context, lend *models.Lend) error {
	return r.DB.Create(lend).Error
}
