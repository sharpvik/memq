package service

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sharpvik/memq/codered"
	"github.com/sharpvik/memq/response"
)

func Health(c echo.Context) error { return c.String(http.StatusOK, "OK") }

func (s *Service) StoreMessage(c echo.Context) error {
	msg, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return codered.NewError(http.StatusBadRequest, err).Respond(c)
	}

	s.queue.Enqueue(msg)

	received := fmt.Sprintf("received: %dB", len(msg))
	log.Println(received)
	return response.New(received).Respond(c)
}
