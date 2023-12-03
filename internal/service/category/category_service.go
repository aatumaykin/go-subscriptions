package category

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
	"git.home/alex/go-subscriptions/internal/domain/category/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

type Service struct {
	repository repository.CategoryRepository
}

type Configuration func(s *Service) error

func NewCategoryService(cfgs ...Configuration) (*Service, error) {
	s := &Service{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func WithCategoryRepository(r repository.CategoryRepository) Configuration {
	return func(cs *Service) error {
		cs.repository = r
		return nil
	}
}

func WithMemoryCategoryRepository() Configuration {
	return WithCategoryRepository(memory.NewCategoryRepository())
}

func (cs *Service) Create(ctx context.Context, name string) (*entity.Category, error) {
	category := entity.Category{
		Name: name,
	}

	return cs.repository.Create(ctx, category)
}

func (cs *Service) Get(ctx context.Context, ID uint) (*entity.Category, error) {
	return cs.repository.Get(ctx, ID)
}

func (cs *Service) GetAll(ctx context.Context) (repository.Categories, error) {
	return cs.repository.GetAll(ctx)
}

func (cs *Service) Update(ctx context.Context, category entity.Category) (*entity.Category, error) {
	return cs.repository.Update(ctx, category)
}

func (cs *Service) Delete(ctx context.Context, ID uint) error {
	return cs.repository.Delete(ctx, ID)
}
