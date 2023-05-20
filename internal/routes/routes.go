package routes

import (
	"net/http"

	"github.com/arpushkarev/max-substring/internal/api"
)

func Route() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/substring/", api.Substring)

	return mux
}
