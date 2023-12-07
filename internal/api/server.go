package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/api/handler/health_handler"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

type HTTPServer struct {
	listenAddr string
	timeout    time.Duration
	router     *httprouter.Router
	ctx        context.Context
}

type Configuration func(s *HTTPServer) error

const defaultTimeout = 5 * time.Second

func NewHTTPServer(cfgs ...Configuration) (*HTTPServer, error) {
	s := &HTTPServer{}

	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func WithListenAddr(addr string) Configuration {
	return func(s *HTTPServer) error {
		s.listenAddr = addr
		return nil
	}
}

func WithTimeout(timeout time.Duration) Configuration {
	return func(s *HTTPServer) error {
		s.timeout = timeout
		return nil
	}
}

func WithRouter(router *httprouter.Router) Configuration {
	return func(s *HTTPServer) error {
		s.router = router
		return nil
	}
}

func WithDefaultRouter() Configuration {
	return WithRouter(httprouter.New())
}

func WithContext(ctx context.Context) Configuration {
	return func(s *HTTPServer) error {
		s.ctx = ctx
		return nil
	}
}

func (s *HTTPServer) ListenAndServe() {
	server := &http.Server{
		Addr:        s.listenAddr,
		Handler:     s.router,
		ReadTimeout: s.timeout * time.Second,
	}

	// Create a channel to receive signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Start the server in a goroutine
	go func() {
		slog.Info("API server started", "address", s.listenAddr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen: %s\n", err)
		}
	}()

	// Wait for a signal
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)

	slog.Info("Shutting down")
}

func WithHealthHandler() Configuration {
	return func(s *HTTPServer) error {
		s.router.GET("/health", health_handler.Handle())
		return nil
	}
}

func WithCategoryCollectionGetHandler(categoryService *service.CategoryService) Configuration {
	return func(s *HTTPServer) error {
		s.router.GET("/api/categories", category_handler.CollectionGetHandle(s.ctx, categoryService))
		return nil
	}
}

func WithCategoryCreateHandler(categoryService *service.CategoryService) Configuration {
	return func(s *HTTPServer) error {
		s.router.POST("/api/category", category_handler.CreateHandle(s.ctx, categoryService))
		return nil
	}
}

func WithCategoryGetHandler(categoryService *service.CategoryService) Configuration {
	return func(s *HTTPServer) error {
		s.router.GET("/api/category/:id", category_handler.GetHandle(s.ctx, categoryService))
		return nil
	}
}

func WithCategoryUpdateHandler(categoryService *service.CategoryService) Configuration {
	return func(s *HTTPServer) error {
		s.router.PUT("/api/category/:id", category_handler.UpdateHandle(s.ctx, categoryService))
		return nil
	}
}

func WithCategoryDeleteHandler(categoryService *service.CategoryService) Configuration {
	return func(s *HTTPServer) error {
		s.router.DELETE("/api/category/:id", category_handler.DeleteHandle(s.ctx, categoryService))
		return nil
	}
}
