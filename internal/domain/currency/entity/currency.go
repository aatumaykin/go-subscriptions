package entity

type Currency struct {
	Code   string
	Symbol string
	Name   string
}

var (
	USD = Currency{Code: "USD", Symbol: "$", Name: "US Dollar"}
	RUB = Currency{Code: "RUB", Symbol: "₽", Name: "Russian Ruble"}
)
