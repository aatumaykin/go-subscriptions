package repository

import (
	"context"
	"errors"

	"git.home/alex/go-subscriptions/internal/domain/entity"
)

var (
	ErrNotFoundCurrency = errors.New("the currency was not found in the repository")
	ErrCreateCurrency   = errors.New("failed to add the currency to the repository")
	ErrUpdateCurrency   = errors.New("failed to update the currency in the repository")
	ErrDeleteCurrency   = errors.New("failed to delete the currency from the repository")
)

type Currencies []entity.Currency

type CurrencyRepository interface {
	Create(ctx context.Context, currency entity.Currency) (*entity.Currency, error)
	Get(ctx context.Context, code string) (*entity.Currency, error)
	GetAll(ctx context.Context) (Currencies, error)
	Update(ctx context.Context, currency entity.Currency) (*entity.Currency, error)
	Delete(ctx context.Context, code string) error
}
