package main

import (
	"context"
	"flag"
	"log"
	"sync"

	"github.com/arpushkarev/max-substring/internal/cli"
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

	c, err := cli.NewCli()
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		err = a.Run()
		if err != nil {
			log.Fatalf("failed to run app %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		c.ResponseWriter()
	}()

	wg.Wait()
}
