.PHONY: run-server
run-server:
	go run -v ./cmd/server

.PHONY: run-cli
run-cli:
	go run -v ./cmd/cli

.PHONY: install-go-deps
install-go-deps:
	ls go.mod || go mod init