package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/sharpvik/memq/codered"
	"github.com/sharpvik/memq/response"
)

func Health(c echo.Context) error { return c.String(http.StatusOK, "OK") }

func (s *Service) SendMessage(c echo.Context) error {
	msg, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return codered.NewError(http.StatusBadRequest, err).Respond(c)
	}

	s.Enqueue(msg)

	return response.New(fmt.Sprintf("received: %dB", len(msg))).Respond(c)
}

func (s *Service) Subscribe(c echo.Context) error {
	rawSubscriberURL := c.Param("url")
	if _, err := url.Parse(rawSubscriberURL); err != nil {
		return codered.NewError(http.StatusBadRequest, err).Respond(c)
	}

	s.SetSubscriber(rawSubscriberURL)

	return response.New(
		fmt.Sprintf("subscriber set: %s", rawSubscriberURL),
	).Respond(c)
}
