build:
	go build -o ./bin/pricefetcher ./cmd

run: build
	./bin/pricefetcher
