package app

import (
	"net/http"

	"github.com/arpushkarev/max-substring/internal/server/api"
)

func Route() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/substring/", api.Substring)

	return mux
}
