package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func Substring(w http.ResponseWriter, r *http.Request) {
	str := r.URL.Query().Get("str")

	bestLen := 0
	letters := map[string]int{}
	start := 0
	var firstLetter, lastLetter int

	for i := 0; i < len(str); i++ {
		if i-start > bestLen {
			bestLen = i - start
			firstLetter = start
			lastLetter = start + (bestLen)
		}

		c := string(str[i])
		if idx, ok := letters[c]; ok && idx >= start {
			start = idx + 1
		}
		letters[c] = i

	}

	substring := str[firstLetter:lastLetter]

	res := strconv.Itoa(bestLen)

	fmt.Fprintf(w, "Max len:%s, String:%s", res, substring)

}
