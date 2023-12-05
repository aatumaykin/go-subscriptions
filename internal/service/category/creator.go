package category

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
	"git.home/alex/go-subscriptions/internal/domain/category/repository"
)

type Creator struct {
	repository.CategoryRepository
}

func NewCreator(r repository.CategoryRepository) *Creator {
	return &Creator{
		CategoryRepository: r,
	}
}

func (c *Creator) Create(ctx context.Context, name string) (*entity.Category, error) {
	category := entity.Category{
		Name: name,
	}

	return c.CategoryRepository.Create(ctx, category)
}
