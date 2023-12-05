package category

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/repository"
)

type Deleter struct {
	repository.CategoryRepository
}

func NewDeleter(r repository.CategoryRepository) *Deleter {
	return &Deleter{
		CategoryRepository: r,
	}
}

func (d *Deleter) Delete(ctx context.Context, ID uint) error {
	return d.CategoryRepository.Delete(ctx, ID)
}
