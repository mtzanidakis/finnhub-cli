.PHONY: build install test lint clean

build:
	go build -o finnhub .

install: build
	install -m 555 finnhub /usr/local/bin/finnhub

test:
	go test ./...

lint:
	golangci-lint run ./...

clean:
	rm -f finnhub
