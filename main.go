package main

import "fmt"

func main() {
	logger := NewLoggerService(&priceFetcher{})
	store, err := NewPostgresStore()
	if err != nil {
		fmt.Print(err)
	}
	store.InitTables()

	server := NewJSONApiServer(":8080", logger, *store)
	server.Run()

}
