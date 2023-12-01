package repository

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
)

type Categories []entity.Category

type CategoryRepository interface {
	Create(ctx context.Context, category entity.Category) (*entity.Category, error)
	GetByID(ctx context.Context, ID uint) (*entity.Category, error)
	GetAll(ctx context.Context) (Categories, error)
	Update(ctx context.Context, category entity.Category) (*entity.Category, error)
	Delete(ctx context.Context, ID uint) error
}
