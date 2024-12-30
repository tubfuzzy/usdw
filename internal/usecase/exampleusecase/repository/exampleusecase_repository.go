package repository

import (
	"context"
	"usdw/config"
	"usdw/internal/domain"
	"usdw/internal/domain/entity"
	"usdw/pkg/db"
	"usdw/pkg/logger"
)

type exampleRepository struct {
}

func NewExampleRepository(db *db.DB, logger logger.Logger, config *config.Configuration) domain.ExampleRepository {

	return &exampleRepository{}
}

func (r exampleRepository) FindUser(ctx context.Context, key string) (entity.User, error) {
	return entity.User{Name: "test-name"}, nil
}
