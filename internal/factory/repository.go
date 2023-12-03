package factory

import (
	categoryrepositoryinterface "git.home/alex/go-subscriptions/internal/domain/category/repository"
	currencyrepositoryinterface "git.home/alex/go-subscriptions/internal/domain/currency/repository"
	cyclerepositoryinterface "git.home/alex/go-subscriptions/internal/domain/cycle/repository"
	subscriptionrepositoryinterface "git.home/alex/go-subscriptions/internal/domain/subscription/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

type RepositoryFactory struct {
	CategoryRepository     categoryrepositoryinterface.CategoryRepository
	CurrencyRepository     currencyrepositoryinterface.CurrencyRepository
	CycleRepository        cyclerepositoryinterface.CycleRepository
	SubscriptionRepository subscriptionrepositoryinterface.SubscriptionRepository
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
