package entity

import (
	"time"

	categoryentity "git.home/alex/go-subscriptions/internal/domain/category/entity"
	currencyentity "git.home/alex/go-subscriptions/internal/domain/currency/entity"
	cycleentity "git.home/alex/go-subscriptions/internal/domain/cycle/entity"
)

type PaymentDate time.Time

type Subscription struct {
	ID              uint
	Name            string
	Note            string
	Logo            string
	Category        categoryentity.Category
	Price           float64
	Currency        currencyentity.Currency
	Cycle           cycleentity.Cycle
	NextPaymentDate PaymentDate
}
