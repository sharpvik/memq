package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/sharpvik/env-go"
	"github.com/sharpvik/memq/service"
)

var (
	sub = env.MustGet("SUB")

	addr = env.GetOr("ADDR", "localhost:5359")
	qcap = env.GetOr("QCAP", "1024")
)

func main() {
	qcapU64, err := strconv.ParseUint(qcap, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	if err := service.
		New(uint(qcapU64), sub).
		Server().
		Start(addr); err != http.ErrServerClosed {
		log.Println(err)
	}
}
