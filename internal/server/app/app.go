package app

import (
	"context"

	"github.com/arpushkarev/max-substring/internal/server/config"
)

type app struct {
	config *config.Config
}

func NewApp(ctx context.Context, path string) (*app, error) {
	cfg, err := config.NewConfig(path)
	if err != nil {
		return nil, err
	}

	return &app{
		config: cfg,
	}, nil
}
