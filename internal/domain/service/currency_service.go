package service

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
)

var (
	ErrInvalidCurrency = errors.New("the currency is not valid")
)

type CurrencyService struct {
	repo repository.CurrencyRepository
}

func NewCurrencyService(repo repository.CurrencyRepository) *CurrencyService {
	return &CurrencyService{
		repo: repo,
	}
}

func (s *CurrencyService) CreateCurrency(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	if currency.Code == "" || currency.Symbol == "" || currency.Name == "" {
		return nil, ErrInvalidCurrency
	}

	return s.repo.Create(ctx, currency)
}

func (s *CurrencyService) GetCurrency(ctx context.Context, code string) (*entity.Currency, error) {
	if code == "" {
		return nil, repository.ErrNotFoundCurrency
	}

	return s.repo.Get(ctx, code)
}

func (s *CurrencyService) GetAllCurrencies(ctx context.Context) (repository.Currencies, error) {
	return s.repo.GetAll(ctx)
}

func (s *CurrencyService) UpdateCurrency(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	if currency.Code == "" || currency.Symbol == "" || currency.Name == "" {
		return nil, ErrInvalidCurrency
	}

	return s.repo.Update(ctx, currency)
}

func (s *CurrencyService) DeleteCurrency(ctx context.Context, code string) error {
	if code == "" {
		return repository.ErrNotFoundCurrency
	}

	return s.repo.Delete(ctx, code)
}
