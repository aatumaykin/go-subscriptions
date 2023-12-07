package subscription_handler

import "git.home/alex/go-subscriptions/internal/domain/service"

type HandlerOpts struct {
	SubscriptionService *service.SubscriptionService
	CategoryService     *service.CategoryService
	CycleService        *service.CycleService
	CurrencyService     *service.CurrencyService
}
