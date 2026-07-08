package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gCastl/go-clean-template/internal/config"
)

type API struct {
	config *config.Config
	srv    *http.Server
}

func (a *API) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.srv.Shutdown(ctx)
}

func (a *API) ListenAndServe() error {
	port := a.config.Port
	mux := http.NewServeMux()

	a.srv = &http.Server{Addr: fmt.Sprintf(":%v", port), Handler: mux}

	slog.Info("API Listening on " + a.srv.Addr)
	return a.srv.ListenAndServe()
}

func NewAPI(c *config.Config) *API {
	return &API{
		config: c,
	}
}
