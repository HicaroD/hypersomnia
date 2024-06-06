OUT=hyper

build:
	go build -o $(OUT) .

run:
	go run .

fmt:
	gofmt -s -w .
