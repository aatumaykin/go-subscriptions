package entity

import (
	"time"
)

type PaymentDate time.Time

type Subscription struct {
	ID    uint
	Price float64
	Category
	Currency
	Cycle
	NextPaymentDate PaymentDate
	Name            string
	Note            string
	Logo            string
}
