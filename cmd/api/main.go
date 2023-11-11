package main

import (
	"log"

	"github.com/mskutle/url-shortener/internal/shortener"
)

func main() {
	store := shortener.NewInMemoryStore()
	server := shortener.NewServer(":3000", store)

	log.Fatal(server.Start())
}
