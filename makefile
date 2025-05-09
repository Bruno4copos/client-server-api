build:
	@go build -o server server.go
	@go build -o client client.go

run-server:
	@go run server.go

run-client:
	@go run client.go

test:
	@go test ./...

docker-build:
	@docker build -t cotacao-app .

docker-run:
	@docker run -p 8080:8080 cotacao-app