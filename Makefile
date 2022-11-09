run:
	go run main.go

build:
	go build -o bin/api main.go

lint:
	golangci-lint run

test:
	go test -v ./...
