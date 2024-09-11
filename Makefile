# Makefile for building and deploying Golang Lambda function with Lambroll

.PHONY: all build deploy clean prepare test rollback

SRC := main.go

BIN := bootstrap

all: test build

build: prepare $(BIN)

$(BIN):
	cd src; GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ../$(BIN) $(SRC)

deploy: clean build
	lambroll deploy --alias="current"

rollback:
	lambroll rollback --alias="current"

clean:
	rm -rf $(BIN)

prepare:
	cd src; go mod tidy

test: prepare
	go test -v ./...

.SHELLFLAGS = -c -e