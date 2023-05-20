package routs

import (
	"fmt"
	"net/http"

	"github.com/arpushkarev/max-substring/internal/api"
	"github.com/go-chi/chi/v5"
)

func Route(url string) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/api/substring", api.Handler)
	fmt.Println(url)

	return mux
}
