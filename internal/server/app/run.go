package app

import "net/http"

func (a *app) Run() error {

	server := http.Server{
		Addr:    a.config.GetHTTPAddress(),
		Handler: Route(),
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
