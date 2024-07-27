package codered

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Error struct {
	status int
	reason any
}

type Response struct {
	Error any `json:"error"`
}

func NewError(status int, reason any) *Error {
	return &Error{
		status: status,
		reason: reason,
	}
}

func (err *Error) Error() string {
	return fmt.Sprintln(err.reason)
}

func (err *Error) Response() *Response {
	return &Response{
		Error: err.reason,
	}
}

func (err *Error) Respond(c echo.Context) error {
	return c.JSON(err.status, err.Response())
}
