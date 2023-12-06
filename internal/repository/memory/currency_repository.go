package memory

import (
	"context"
	"sync"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
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

func (r *CurrencyRepository) Create(_ context.Context, currency entity.Currency) (*entity.Currency, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.currencies[currency.Code]; ok {
		return nil, repository.ErrAlreadyExistsCurrency
	}

	r.currencies[currency.Code] = currency

	return &currency, nil
}

func (r *CurrencyRepository) Get(_ context.Context, code string) (*entity.Currency, error) {
	r.Lock()
	defer r.Unlock()

	if currency, ok := r.currencies[code]; ok {
		return &currency, nil
	}

	return nil, repository.ErrNotFoundCurrency
}

func (r *CurrencyRepository) GetAll(_ context.Context) (repository.Currencies, error) {
	r.Lock()
	defer r.Unlock()

	var currencies repository.Currencies

	for _, currency := range r.currencies {
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (r *CurrencyRepository) Update(_ context.Context, currency entity.Currency) (*entity.Currency, error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.currencies[currency.Code]; !ok {
		return nil, repository.ErrNotFoundCurrency
	}

	r.currencies[currency.Code] = currency

	return &currency, nil
}

func (r *CurrencyRepository) Delete(_ context.Context, code string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.currencies[code]; !ok {
		return repository.ErrNotFoundCurrency
	}

	delete(r.currencies, code)

	return nil
}
