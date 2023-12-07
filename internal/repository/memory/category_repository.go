package memory

import (
	"context"
	"sync"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
)

type CategoryRepository struct {
	categories map[uint]entity.Category
	sync.Mutex
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		categories: make(map[uint]entity.Category),
	}
}

func (r *CategoryRepository) Create(_ context.Context, category entity.Category) (*entity.Category, error) {
	r.Lock()
	defer r.Unlock()

	category.ID = uint(len(r.categories) + 1)
	r.categories[category.ID] = category

	return &category, nil
}

func (r *CategoryRepository) Get(_ context.Context, id uint) (*entity.Category, error) {
	r.Lock()
	defer r.Unlock()

	category, ok := r.categories[id]
	if !ok {
		return nil, repository.ErrNotFoundCategory
	}

	return &category, nil
}

func (r *CategoryRepository) GetAll(_ context.Context) (repository.Categories, error) {
	r.Lock()
	defer r.Unlock()

	var categories repository.Categories
	for _, category := range r.categories {
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) Update(_ context.Context, category entity.Category) (*entity.Category, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.categories[category.ID]; !ok {
		return nil, repository.ErrNotFoundCategory
	}

	r.categories[category.ID] = category

	return &category, nil
}

func (r *CategoryRepository) Delete(_ context.Context, id uint) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.categories[id]; !ok {
		return repository.ErrNotFoundCategory
	}

	delete(r.categories, id)

	return nil
}
