package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	producer = "localhost:8000"
	consumer = "localhost:9000"
	memq     = "localhost:5359"
)

func main() {
	e := echo.New()

	e.POST("/msg", func(c echo.Context) error {
		msg, err := io.ReadAll(c.Request().Body)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, err.Error())
		}

		time.Sleep(5 * time.Second)
		log.Printf("message read: %dB", len(msg))

		return c.String(http.StatusOK, "OK")
	})

	if err := e.Start(consumer); err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
