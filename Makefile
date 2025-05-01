build:
	go build -o ./bin/pricefetcher ./cmd/api

run: build
	./bin/pricefetcher
