package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*Service
	*echo.Echo
}

func (s *Server) Start(addr string) error {
	go s.Service.ForwardMessages()
	return s.Echo.Start(addr)
}

func (s *Service) Server() *Server {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/health", Health)
	e.POST("/msg", s.StoreMessage)

	return &Server{
		Service: s,
		Echo:    e,
	}
}
