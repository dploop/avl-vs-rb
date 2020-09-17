MAIN_ROOT ?= ./cmd/avl-vs-rb

build:
	cd $(MAIN_ROOT) && go build

run:
	cd $(MAIN_ROOT) && go run .

lint:
	golangci-lint run
