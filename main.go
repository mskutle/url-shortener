package main

import (
	"log"
)

func main() {
	store := NewInMemoryUrlStore()
	server := NewServer(":3000", store)

	log.Fatal(server.Start())
}
