package factory

import (
	"git.home/alex/go-subscriptions/internal/domain/service"
)

type ServiceFactory struct {
	repositoryFactory   *RepositoryFactory
	CategoryService     *service.CategoryService
	CurrencyService     *service.CurrencyService
	CycleService        *service.CycleService
	SubscriptionService *service.SubscriptionService
}

type ServiceConfiguration func(sf *ServiceFactory) error

func NewServiceFactory(cfgs ...ServiceConfiguration) (*ServiceFactory, error) {
	sf := &ServiceFactory{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(sf)
		if err != nil {
			return nil, err
		}
	}

	return sf, nil
}

func WithRepositoryFactory(repositoryFactory *RepositoryFactory) ServiceConfiguration {
	return func(sf *ServiceFactory) error {
		sf.repositoryFactory = repositoryFactory
		return nil
	}
}

func WithCategoryService() ServiceConfiguration {
	return func(sf *ServiceFactory) error {
		sf.CategoryService = service.NewCategoryService(sf.repositoryFactory.CategoryRepository)
		return nil
	}
}

func WithCurrencyService() ServiceConfiguration {
	return func(sf *ServiceFactory) error {
		sf.CurrencyService = service.NewCurrencyService(sf.repositoryFactory.CurrencyRepository)
		return nil
	}
}

func WithCycleService() ServiceConfiguration {
	return func(sf *ServiceFactory) error {
		sf.CycleService = service.NewCycleService(sf.repositoryFactory.CycleRepository)
		return nil
	}
}

func WithSubscriptionService() ServiceConfiguration {
	return func(sf *ServiceFactory) error {
		sf.SubscriptionService = service.NewSubscriptionService(sf.repositoryFactory.SubscriptionRepository)
		return nil
	}
}
