package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cache struct {
	URL string
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
	list[0], list[1] = list[1], list[0]
	url := strings.Join(list, "-")

	return &Cache{
		URL: url,
	}, nil
}
