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

func (s *Service) Server(auth *Auth) *Server {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/health", Health)
	e.POST("/msg", s.StoreMessage, auth.Middleware())

	return &Server{
		Service: s,
		Echo:    e,
	}
}
