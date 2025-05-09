FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod init cotacao && \\
    go get github.com/mattn/go-sqlite3 && \\
    go build -o server server.go && \\
    go build -o client client.go

EXPOSE 8080

CMD ["./server"]