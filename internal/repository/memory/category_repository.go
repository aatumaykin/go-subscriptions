package memory

import (
	"context"
	"fmt"
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

func (r *CategoryRepository) Create(ctx context.Context, category entity.Category) (*entity.Category, error) {
	r.Lock()
	defer r.Unlock()

	category.ID = uint(len(r.categories) + 1)
	r.categories[category.ID] = category

	return &category, nil
}

func (r *CategoryRepository) Get(ctx context.Context, ID uint) (*entity.Category, error) {
	r.Lock()
	defer r.Unlock()

	category, ok := r.categories[ID]
	if !ok {
		return nil, fmt.Errorf("category not found: %w", repository.ErrNotFoundCategory)
	}

	return &category, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) (repository.Categories, error) {
	r.Lock()
	defer r.Unlock()

	var categories repository.Categories
	for _, category := range r.categories {
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) Update(ctx context.Context, category entity.Category) (*entity.Category, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.categories[category.ID]; !ok {
		return nil, fmt.Errorf("category not found: %w", repository.ErrUpdateCategory)
	}

	r.categories[category.ID] = category

	return &category, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, ID uint) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.categories[ID]; !ok {
		return fmt.Errorf("category not found: %w", repository.ErrDeleteCategory)
	}

	delete(r.categories, ID)

	return nil
}
