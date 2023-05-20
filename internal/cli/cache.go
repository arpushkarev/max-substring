package cli

import (
	"bufio"
	"fmt"
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

	return &Cache{
		url: url,
	}, nil
}

func (c *Cache) GetURL() string {
	return c.url
}
