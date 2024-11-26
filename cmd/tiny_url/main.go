package main

import (
	"log"

	"github.com/edjubert/tiny-url/internal/api"
)

func main() {
	server, err := api.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
