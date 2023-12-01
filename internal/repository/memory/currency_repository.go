package memory

import (
	"context"
	"fmt"
	"sync"

	"git.home/alex/go-subscriptions/internal/domain/currency/entity"
	"git.home/alex/go-subscriptions/internal/domain/currency/repository"
)

type CurrencyRepository struct {
	currencies map[string]entity.Currency
	sync.Mutex
}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{
		currencies: make(map[string]entity.Currency),
	}
}

func (r *CurrencyRepository) Create(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.currencies[currency.Code]; ok {
		return nil, fmt.Errorf("currency already exists: %w", repository.ErrCreateCurrency)
	}

	r.currencies[currency.Code] = currency

	return &currency, nil
}

func (r *CurrencyRepository) Get(ctx context.Context, code string) (*entity.Currency, error) {
	r.Lock()
	defer r.Unlock()

	if currency, ok := r.currencies[code]; ok {
		return &currency, nil
	}

	return nil, repository.ErrNotFoundCurrency
}

func (r *CurrencyRepository) GetAll(ctx context.Context) (repository.Currencies, error) {
	r.Lock()
	defer r.Unlock()

	var currencies repository.Currencies

	for _, currency := range r.currencies {
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (r *CurrencyRepository) Update(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.currencies[currency.Code]; !ok {
		return nil, fmt.Errorf("currency does not exist: %w", repository.ErrUpdateCurrency)
	}

	r.currencies[currency.Code] = currency

	return &currency, nil
}

func (r *CurrencyRepository) Delete(ctx context.Context, code string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.currencies[code]; !ok {
		return fmt.Errorf("currency does not exist: %w", repository.ErrDeleteCurrency)
	}

	delete(r.currencies, code)

	return nil
}
