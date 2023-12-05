package category

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
	"git.home/alex/go-subscriptions/internal/domain/category/repository"
)

type Updater struct {
	repository.CategoryRepository
}

func NewUpdater(r repository.CategoryRepository) *Updater {
	return &Updater{
		CategoryRepository: r,
	}
}

func (u *Updater) Update(ctx context.Context, category entity.Category) (*entity.Category, error) {
	return u.CategoryRepository.Update(ctx, category)
}
