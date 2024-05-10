
# Makefile for Go build and test
.PHONY: all build run-files run-stdin test docker docker-run docker-stdin docker-test clean

all: build

build:
	go build -o trigrams ./cmd/main.go

run-files:
	go run ./cmd/main.go $(ARGS)

run-stdin:
	cat ./texts/*.txt | go run ./cmd/main.go

test:
	go test ./...

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -f trigrams


# Docker commands
docker:
	docker build -t trigrams:latest .

docker-run:
	docker run --rm -v "$$(pwd)/texts:/texts" trigrams:latest -- /texts/moby_dick.txt /texts/brothers_karamazov.txt

docker-stdin:
	cat ./texts/*.txt | docker run --rm -i trigrams:latest

docker-test:
	docker build -f Dockerfile.test -t trigrams-test:latest .
	docker run --rm trigrams-test:latest
