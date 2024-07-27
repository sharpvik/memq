package main

import (
	"log"
	"net/http"

	"github.com/sharpvik/env-go"
	"github.com/sharpvik/memq/service"
)

var address = env.GetOr("ADDRESS", "localhost:5359")

func main() {
	if err := service.Server().Start(address); err != http.ErrServerClosed {
		log.Println(err)
	}
}
