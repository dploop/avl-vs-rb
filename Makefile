.PHONY: main

main:
	go build -o avl-vs-rb ./main

lint:
	golangci-lint run
