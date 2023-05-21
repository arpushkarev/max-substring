package main

import (
	"log"

	"github.com/arpushkarev/max-substring/internal/cli"
)

func main() {

	for {
		c, err := cli.NewCli()
		if err != nil {
			log.Fatal(err)
		}

		c.ResponseWriter()
	}
}
