package cli

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Cli struct {
	URL string
}

func NewCli() (*Cli, error) {
	fmt.Println("Введите запрос:")

	buffer := bufio.NewReader(os.Stdin)
	input, err := buffer.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.TrimRight(input, "\n \r")
	list := strings.Fields(input)

	if len(list) != 2 {
		return nil, err
	}

	reqUrl := list[1]
	str := list[0]

	if reqUrl != "api/substring" {
		return nil, err
	}

	url := fmt.Sprintf("%s/?str=%s", reqUrl, str)

	return &Cli{
		URL: url,
	}, nil
}

func (c *Cli) ResponseWriter() {

	url := c.URL

	if !strings.HasPrefix(c.URL, "http://localhost:8090/") && !strings.HasPrefix(c.URL, "https://localhost:8090/") {
		url = "http://localhost:8090/" + c.URL
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
	}

	fmt.Println(string(body))

}
