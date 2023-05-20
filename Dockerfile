FROM golang:1.20
WORKDIR /app
COPY go.mod go.sum ./
COPY go_build_main_go ./
EXPOSE 8090