package cli

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *cli) ResponseWriter() error {

	url := c.url

	if !strings.HasPrefix(c.url, "http://localhost:8090/") && !strings.HasPrefix(c.url, "https://localhost:8090/") {
		url = "http://localhost:8090/" + c.url
	}

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
