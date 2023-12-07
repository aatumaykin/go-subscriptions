package api

import (
	"git.home/alex/go-subscriptions/internal/api/handler"
	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/api/handler/cycle_handler"
	"git.home/alex/go-subscriptions/internal/api/handler/empty_handler"
	"git.home/alex/go-subscriptions/internal/api/handler/health_handler"
	"git.home/alex/go-subscriptions/internal/domain/service"
)

func WithHealthHandler() Configuration {
	return func(s *HTTPServer) error {
		s.router.GET("/health", health_handler.Handle())
		return nil
	}
}

func WithCategoryHandlers(cs *service.CategoryService) Configuration {
	return func(s *HTTPServer) error {
		s.router.POST("/api/category", handler.Handle(category_handler.CreateCategory(s.ctx, cs)))
		s.router.GET("/api/category/:id", handler.Handle(category_handler.GetCategory(s.ctx, cs)))
		s.router.GET("/api/categories", handler.Handle(category_handler.GetCategories(s.ctx, cs)))
		s.router.PUT("/api/category/:id", handler.Handle(category_handler.UpdateCategory(s.ctx, cs)))
		s.router.DELETE("/api/category/:id", handler.Handle(category_handler.DeleteCategory(s.ctx, cs)))

		return nil
	}
}

func WithCurrencyHandlers(cs *service.CurrencyService) Configuration {
	return func(s *HTTPServer) error {
		s.router.POST("/api/currency", handler.Handle(currency_handler.CreateCurrency(s.ctx, cs)))
		s.router.GET("/api/currency/:code", handler.Handle(currency_handler.GetCurrency(s.ctx, cs)))
		s.router.GET("/api/currencies", handler.Handle(currency_handler.GetCurrencies(s.ctx, cs)))
		s.router.PUT("/api/currency/:code", handler.Handle(currency_handler.UpdateCurrency(s.ctx, cs)))
		s.router.DELETE("/api/currency/:code", handler.Handle(currency_handler.DeleteCurrency(s.ctx, cs)))

		return nil
	}
}

func WithCycleHandlers(cs *service.CycleService) Configuration {
	return func(s *HTTPServer) error {
		s.router.POST("/api/cycle", handler.Handle(cycle_handler.CreateCycle(s.ctx, cs)))
		s.router.GET("/api/cycle/:id", handler.Handle(cycle_handler.GetCycle(s.ctx, cs)))
		s.router.GET("/api/cycles", handler.Handle(cycle_handler.GetCycles(s.ctx, cs)))
		s.router.PUT("/api/cycle/:id", handler.Handle(cycle_handler.UpdateCycle(s.ctx, cs)))
		s.router.DELETE("/api/cycle/:id", handler.Handle(cycle_handler.DeleteCycle(s.ctx, cs)))

		return nil
	}
}

func WithSubscribeHandlers(_ *service.SubscriptionService) Configuration {
	return func(s *HTTPServer) error {
		s.router.POST("/api/subscription", empty_handler.Handle())
		s.router.GET("/api/subscription/:id", empty_handler.Handle())
		s.router.GET("/api/subscriptions", empty_handler.Handle())
		s.router.PUT("/api/subscription/:id", empty_handler.Handle())
		s.router.DELETE("/api/subscription/:id", empty_handler.Handle())

		return nil
	}
}
