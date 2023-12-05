package category

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/repository"
)

type CollectionGetter struct {
	repository.CategoryRepository
}

func NewCollectionGetter(r repository.CategoryRepository) *CollectionGetter {
	return &CollectionGetter{
		CategoryRepository: r,
	}
}

func (cg *CollectionGetter) GetAll(ctx context.Context) (repository.Categories, error) {
	return cg.CategoryRepository.GetAll(ctx)
}
