package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cli struct {
	url string
}

func NewCli() (*cli, error) {
	fmt.Println("(чтобы остановить приложение введите \"stop\") \n Введите запрос:")

	buffer := bufio.NewReader(os.Stdin)
	input, err := buffer.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.TrimRight(input, "\n \r")

	if input == "stop" {
		err = errors.New("приложение остановлено пользователем")
		return nil, err
	}

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

	return &cli{
		url: url,
	}, nil
}
