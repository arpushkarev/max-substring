FROM golang:1.20-alpine
WORKDIR /app
COPY ./ /app
EXPOSE 8090
CMD ["./server"]