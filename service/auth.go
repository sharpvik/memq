package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Auth struct {
	key string
}

func NewAuth(key string) *Auth {
	return &Auth{
		key: key,
	}
}

func (auth *Auth) Validate(key string, c echo.Context) (bool, error) {
	return key == auth.key, nil
}

func (auth *Auth) Middleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(auth.Validate)
}
