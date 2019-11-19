.PHONY: main

main:
	go build -o avl-vs-rb ./main

test: test_avl test_rb

test_avl:
	go test -v ./avl/...

test_rb:
	go test -v ./rb/...

bench: bench_avl bench_rb

bench_avl:
	go test -v -bench . -run ^$$ ./avl/...

bench_rb:
	go test -v -bench . -run ^$$ ./rb/...

lint:
	golangci-lint run
