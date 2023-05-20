package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Substring(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()

	list := strings.Split(url, "-")

	str := list[1]

	bestLen := 0

	letters := map[string]int{}
	start := 0
	for i := 0; ; i++ {
		if i-start > bestLen {
			bestLen = i - start
		}

		if i >= len(str) {
			break
		}

		c := string(str[i])
		if idx, ok := letters[c]; ok && idx >= start {
			start = idx + 1
		}
		letters[c] = i
	}

	res := strconv.Itoa(bestLen)

	fmt.Fprintf(w, res)

}
