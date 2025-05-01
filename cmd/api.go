package main

import (
	"fmt"

	"github.com/ebhlz88/est-shop/server"
	"github.com/ebhlz88/est-shop/store"
)

func main() {
	// logger := NewLoggerService(&priceFetcher{})
	store, err := store.NewPostgresStore()
	if err != nil {
		fmt.Print(err)
	}
	store.InitTables()

	svc := server.NewJSONApiServer(":8080", store)
	server.Run(svc)

}
