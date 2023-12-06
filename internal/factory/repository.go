package factory

import (
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

type RepositoryFactory struct {
	repository.CategoryRepository
	repository.CurrencyRepository
	repository.CycleRepository
	repository.SubscriptionRepository
}

type RepositoryConfiguration func(rf *RepositoryFactory) error

func NewRepositoryFactory(cfgs ...RepositoryConfiguration) (*RepositoryFactory, error) {
	f := &RepositoryFactory{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(f)
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func WithMemoryRepository() RepositoryConfiguration {
	return func(rf *RepositoryFactory) error {
		rf.CategoryRepository = memory.NewCategoryRepository()
		rf.CurrencyRepository = memory.NewCurrencyRepository()
		rf.CycleRepository = memory.NewCycleRepository()
		rf.SubscriptionRepository = memory.NewSubscriptionRepository()
		return nil
	}
}

func WithRedisRepository() RepositoryConfiguration {
	return func(rf *RepositoryFactory) error {
		return ErrNotImplemented
	}
}

func WithSqliteRepository() RepositoryConfiguration {
	return func(rf *RepositoryFactory) error {
		return ErrNotImplemented
	}
}
