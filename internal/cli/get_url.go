package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func (c *cli) GetURL() error {
	fmt.Println("(чтобы остановить приложение введите \"stop\") \n Введите запрос:")

	buffer := bufio.NewReader(os.Stdin)
	input, err := buffer.ReadString('\n')
	if err != nil {
		return err
	}

	input = strings.TrimRight(input, "\n \r")

	if input == "stop" {
		err = errors.New("приложение остановлено пользователем")
		return err
	}

	list := strings.Fields(input)

	if len(list) != 2 {
		return err
	}

	reqUrl := list[1]
	str := list[0]

	if reqUrl != "api/substring" {
		return err
	}

	url := fmt.Sprintf("%s/?str=%s", reqUrl, str)

	c.url = url

	return nil
}
