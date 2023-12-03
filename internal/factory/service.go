package factory

import (
	categoryserviceinterface "git.home/alex/go-subscriptions/internal/domain/category/service"
	currencyserviceinterface "git.home/alex/go-subscriptions/internal/domain/currency/service"
	cycleserviceinterface "git.home/alex/go-subscriptions/internal/domain/cycle/service"
	subscriptionserviceinterface "git.home/alex/go-subscriptions/internal/domain/subscription/service"
	"git.home/alex/go-subscriptions/internal/service/category"
	"git.home/alex/go-subscriptions/internal/service/currency"
	"git.home/alex/go-subscriptions/internal/service/cycle"
	"git.home/alex/go-subscriptions/internal/service/subscription"
)

type ServiceFactory struct {
	repositoryFactory   *RepositoryFactory
	CategoryService     categoryserviceinterface.CategoryService
	CurrencyService     currencyserviceinterface.CurrencyService
	CycleService        cycleserviceinterface.CycleService
	SubscriptionService subscriptionserviceinterface.SubscriptionService
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
		cs, err := category.NewCategoryService(
			category.WithCategoryRepository(sf.repositoryFactory.CategoryRepository),
		)
		if err != nil {
			return err
		}

		sf.CategoryService = cs
		return nil
	}
}

func WithCurrencyService() ServiceConfiguration {
	return func(sf *ServiceFactory) error {
		cs, err := currency.NewCurrencyService(
			currency.WithCurrencyRepository(sf.repositoryFactory.CurrencyRepository),
		)
		if err != nil {
			return err
		}

		sf.CurrencyService = cs
		return nil
	}
}

func WithCycleService() ServiceConfiguration {
	return func(sf *ServiceFactory) error {
		cs, err := cycle.NewCycleService(
			cycle.WithCycleRepository(sf.repositoryFactory.CycleRepository),
		)
		if err != nil {
			return err
		}

		sf.CycleService = cs
		return nil
	}
}

func WithSubscriptionService() ServiceConfiguration {
	return func(sf *ServiceFactory) error {
		ss, err := subscription.NewSubscriptionService(
			subscription.WithSubscriptionRepository(sf.repositoryFactory.SubscriptionRepository),
		)
		if err != nil {
			return err
		}

		sf.SubscriptionService = ss

		return nil
	}
}
