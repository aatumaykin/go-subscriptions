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

type Configuration func(s *Service) error

func NewCurrencyService(cfgs ...Configuration) (*Service, error) {
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

func WithCurrencyRepository(r repository.CurrencyRepository) Configuration {
	return func(s *Service) error {
		s.repository = r
		return nil
	}
}

func WithMemoryCurrencyRepository() Configuration {
	return WithCurrencyRepository(memory.NewCurrencyRepository())
}

func (s *Service) Create(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	return s.repository.Create(ctx, currency)
}

func (s *Service) Get(ctx context.Context, code string) (*entity.Currency, error) {
	return s.repository.Get(ctx, code)
}

func (s *Service) GetAll(ctx context.Context) (repository.Currencies, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) Update(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	return s.repository.Update(ctx, currency)
}

func (s *Service) Delete(ctx context.Context, code string) error {
	return s.repository.Delete(ctx, code)
}
