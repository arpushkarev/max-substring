package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Cache struct {
	url string
}

func NewCache() (*Cache, error) {

	fmt.Println("Введите запрос:")

	buffer := bufio.NewReader(os.Stdin)
	input, err := buffer.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.TrimRight(input, "\n \r")
	list := strings.Split(input, " ")
	if len(list) != 2 {
		log.Fatalf("Incorrect input data")
	}

	//list[0], list[1] = list[1], list[0]
	//url := strings.Join(list, "-")
	reqUrl := list[1]
	str := list[0]

	if reqUrl != "/api/substring" {
		log.Println("Incorrect URL")
	}

	url := fmt.Sprintf("%s/?str=%s", reqUrl, str)
	fmt.Println(url)
	return &Cache{
		url: url,
	}, nil
}

func (c *Cache) GetURL() string {
	return c.url
}
