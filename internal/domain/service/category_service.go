package service

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
)

var (
	ErrInvalidCategory = errors.New("the category is not valid")
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, category entity.Category) (*entity.Category, error) {
	if category.Name == "" {
		return nil, ErrInvalidCategory
	}

	return s.repo.Create(ctx, category)
}

func (s *CategoryService) GetCategory(ctx context.Context, ID uint) (*entity.Category, error) {
	if ID == 0 {
		return nil, repository.ErrNotFoundCategory
	}

	return s.repo.Get(ctx, ID)
}

func (s *CategoryService) GetAllCategories(ctx context.Context) (repository.Categories, error) {
	return s.repo.GetAll(ctx)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category entity.Category) (*entity.Category, error) {
	if category.ID == 0 || category.Name == "" {
		return nil, ErrInvalidCategory
	}

	return s.repo.Update(ctx, category)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, ID uint) error {
	if ID == 0 {
		return repository.ErrNotFoundCategory
	}

	return s.repo.Delete(ctx, ID)
}
