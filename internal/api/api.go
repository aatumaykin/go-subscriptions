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

	"git.home/alex/go-subscriptions/internal/api/handler"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	ListenAddr string
	Timeout    time.Duration
}

func NewAPI(listenAddr string, timeout time.Duration) *API {
	return &API{
		ListenAddr: listenAddr,
		Timeout:    timeout,
	}
}

func (a *API) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	router.GET("/health", handler.Health())

	return router
}

func (a *API) ListenAndServe() {
	router := a.InitRoutes()

	server := &http.Server{
		Addr:        a.ListenAddr,
		Handler:     router,
		ReadTimeout: a.Timeout * time.Second,
	}

	// Create a channel to receive signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Start the server in a goroutine
	go func() {
		slog.Info("API server started", "address", a.ListenAddr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen: %s\n", err)
		}
	}()

	// Wait for a signal
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)

	slog.Info("Shutting down")
}
