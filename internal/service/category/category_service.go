package category

import (
	"git.home/alex/go-subscriptions/internal/domain/category/repository"
	"git.home/alex/go-subscriptions/internal/domain/category/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

type Service struct {
	repository.CategoryRepository
	service.Getter
	service.CollectionGetter
	service.Creator
	service.Updater
	service.Deleter
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
		cs.CategoryRepository = r
		return nil
	}
}

func WithMemoryCategoryRepository() Configuration {
	return WithCategoryRepository(memory.NewCategoryRepository())
}

func WithCollectionGetter() Configuration {
	return func(cs *Service) error {
		cs.CollectionGetter = NewCollectionGetter(cs.CategoryRepository)
		return nil
	}
}

func WithGetter() Configuration {
	return func(cs *Service) error {
		cs.Getter = NewGetter(cs.CategoryRepository)
		return nil
	}
}

func WithUpdater() Configuration {
	return func(cs *Service) error {
		cs.Updater = NewUpdater(cs.CategoryRepository)
		return nil
	}
}

func WithCreator() Configuration {
	return func(cs *Service) error {
		cs.Creator = NewCreator(cs.CategoryRepository)
		return nil
	}
}

func WithDeleter() Configuration {
	return func(cs *Service) error {
		cs.Deleter = NewDeleter(cs.CategoryRepository)
		return nil
	}
}
