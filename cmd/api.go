package main

import (
	"fmt"
	"log"

	"github.com/ebhlz88/est-shop/server"
	"github.com/ebhlz88/est-shop/store"
	"github.com/joho/godotenv"
)

func main() {
	// logger := NewLoggerService(&priceFetcher{})
	err := godotenv.Load("development.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	store, err := store.NewPostgresStore()
	if err != nil {
		fmt.Print(err)
	}
	if err := store.InitTables(); err != nil {
		fmt.Println(err)
	}

	svc := server.NewJSONApiServer(":8080", store)
	server.Run(svc)

}
