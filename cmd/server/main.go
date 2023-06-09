package main

import (
	"context"
	"flag"
	"log"

	"github.com/arpushkarev/max-substring/internal/server/app"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "config/config.json", "Path to configuration")
}

func main() {

	flag.Parse()

	ctx := context.Background()

	a, err := app.NewApp(ctx, pathConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %s", err.Error())
	}
}
