package service

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/currency/entity"
	"git.home/alex/go-subscriptions/internal/domain/currency/repository"
)

type CurrencyService interface {
	Create(ctx context.Context, currency entity.Currency) (*entity.Currency, error)
	Get(ctx context.Context, code string) (*entity.Currency, error)
	GetAll(ctx context.Context) (repository.Currencies, error)
	Update(ctx context.Context, currency entity.Currency) (*entity.Currency, error)
	Delete(ctx context.Context, code string) error
}
