package service

import (
	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	e := echo.New()
	s := New()

	e.GET("/health", Health)
	e.POST("/msg", s.SendMessage)
	e.PUT("/sub/:url", s.Subscribe)

	return e
}
