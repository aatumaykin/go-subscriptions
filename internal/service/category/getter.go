package category

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
	"git.home/alex/go-subscriptions/internal/domain/category/repository"
)

type Getter struct {
	repository.CategoryRepository
}

func NewGetter(r repository.CategoryRepository) *Getter {
	return &Getter{
		CategoryRepository: r,
	}
}

func (g *Getter) Get(ctx context.Context, ID uint) (*entity.Category, error) {
	return g.CategoryRepository.Get(ctx, ID)
}
