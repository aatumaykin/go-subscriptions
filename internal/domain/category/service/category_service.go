package service

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
	"git.home/alex/go-subscriptions/internal/domain/category/repository"
)

type CategoryService interface {
	Create(ctx context.Context, name string) (*entity.Category, error)
	Get(ctx context.Context, ID uint) (*entity.Category, error)
	GetAll(ctx context.Context) (repository.Categories, error)
	Update(ctx context.Context, category entity.Category) (*entity.Category, error)
	Delete(ctx context.Context, ID uint) error
}
