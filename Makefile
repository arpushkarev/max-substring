.PHONY: run
run:
	go run -v ./cmd/server

.PHONY: install-go-deps
install-go-deps:
	ls go.mod || go mod init
	go get github.com/go-chi/chi/v5