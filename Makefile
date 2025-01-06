OUT=hyper

.PHONY: build
build:
	go build -o $(OUT) .

.PHONY: run
run:
	go run .

.PHONY: fmt
fmt:
	gofmt -s -w .
