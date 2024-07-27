package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	memq = "http://localhost:5359/msg"
)

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 5)

		resp, err := http.Post(memq, "application/octet-stream", strings.NewReader("hello"))
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("response status:", resp.Status)
	}
}
