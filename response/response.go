package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Result any `json:"result"`
}

func New(result any) *Response {
	return &Response{
		Result: result,
	}
}

func (resp *Response) Respond(c echo.Context) error {
	return c.JSON(http.StatusOK, resp)
}
