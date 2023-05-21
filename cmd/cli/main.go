package main

import (
	"log"

	"github.com/arpushkarev/max-substring/internal/cli"
)

func main() {

	c, err := cli.NewCli()
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = c.GetURL()
		if err != nil {
			log.Fatal(err)
		}

		err = c.ResponseWriter()
		if err != nil {
			log.Fatal(err)
		}
	}
}
