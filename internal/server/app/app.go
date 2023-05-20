package app

import (
	"context"
	"net/http"

	"github.com/arpushkarev/max-substring/internal/server/config"
)

type App struct {
	config *config.Config
}

func NewApp(ctx context.Context, path string) (*App, error) {
	cfg, err := config.NewConfig(path)
	if err != nil {
		return nil, err
	}

	return &App{
		config: cfg,
	}, nil
}

func (a *App) Run() error {

	server := http.Server{
		Addr:    a.config.GetHTTPAddress(),
		Handler: Route(),
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
