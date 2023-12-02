package repository

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
)

var (
	ErrNotFoundCategory = errors.New("the category was not found in the repository")
	ErrCreateCategory   = errors.New("failed to add the category to the repository")
	ErrUpdateCategory   = errors.New("failed to update the category in the repository")
	ErrDeleteCategory   = errors.New("failed to delete the category from the repository")
)

type Categories []entity.Category

type CategoryRepository interface {
	Create(ctx context.Context, category entity.Category) (*entity.Category, error)
	GetByID(ctx context.Context, ID uint) (*entity.Category, error)
	GetAll(ctx context.Context) (Categories, error)
	Update(ctx context.Context, category entity.Category) (*entity.Category, error)
	Delete(ctx context.Context, ID uint) error
}
