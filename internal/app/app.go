package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/arpushkarev/max-substring/internal/cli"
	"github.com/arpushkarev/max-substring/internal/config"
	"github.com/arpushkarev/max-substring/internal/routs"
)

type App struct {
	Config *config.Config
	ctx    context.Context
	cache  *cli.Cache
}

func NewApp(ctx context.Context, path string) (*App, error) {
	cfg, err := config.NewConfig(path)
	if err != nil {
		return nil, err
	}

	cache, err := cli.NewCache()

	return &App{
		Config: cfg,
		ctx:    ctx,
		cache:  cache,
	}, nil
}

func (a *App) Run() error {

	server := http.Server{
		Addr:    a.Config.GetHTTPAddress(),
		Handler: routs.Route(a.cache.URL),
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) ResponseWriter() {

	url := a.cache.URL

	if !strings.HasPrefix(a.cache.URL, "http://localhost:8090/") && !strings.HasPrefix(a.cache.URL, "https://localhost:8090/") {
		url = "http://localhost:8090/" + a.cache.URL
	}

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to make the HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
		return
	}

	fmt.Printf("The max substring len: %s", string(body))

}
