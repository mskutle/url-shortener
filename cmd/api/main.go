package main

import (
	"log"
	"os"

	"github.com/mskutle/url-shortener/internal/shortener"
)

func main() {
	// store := shortener.NewInMemoryStore()
	store := shortener.NewRedisStore(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"))
	server := shortener.NewServer(store)

	log.Fatal(server.Start())
}
