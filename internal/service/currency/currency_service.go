package currency

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/currency/entity"
	"git.home/alex/go-subscriptions/internal/domain/currency/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

type Service struct {
	repository repository.CurrencyRepository
}

type ServiceConfiguration func(cs *Service) error

func NewCurrencyService(cfgs ...ServiceConfiguration) (*Service, error) {
	cs := &Service{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(cs)
		if err != nil {
			return nil, err
		}
	}

	return cs, nil
}

func WithCurrencyRepository(r repository.CurrencyRepository) ServiceConfiguration {
	return func(cs *Service) error {
		cs.repository = r
		return nil
	}
}

func WithMemoryCurrencyRepository() ServiceConfiguration {
	return WithCurrencyRepository(memory.NewCurrencyRepository())
}

func (cs *Service) Create(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	return cs.repository.Create(ctx, currency)
}

func (cs *Service) Get(ctx context.Context, code string) (*entity.Currency, error) {
	return cs.repository.Get(ctx, code)
}

func (cs *Service) GetAll(ctx context.Context) (repository.Currencies, error) {
	return cs.repository.GetAll(ctx)
}

func (cs *Service) Update(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	return cs.repository.Update(ctx, currency)
}

func (cs *Service) Delete(ctx context.Context, code string) error {
	return cs.repository.Delete(ctx, code)
}
