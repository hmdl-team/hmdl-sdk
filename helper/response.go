package helper

import (
	"github.com/labstack/echo/v4"

	"net/http"
)

type Response struct {
	StatusCode int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ResponseWithCode(c echo.Context, code int, errMsg ...string) error {
	var msg string
	if len(errMsg) == 0 {
		msg = http.StatusText(code)
	} else {
		msg = errMsg[0]
	}
	return c.JSON(code, Response{
		StatusCode: code,
		Message:    msg,
	})
}

func ResponseDataMessage(c echo.Context, mesage string, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    mesage,
		Data:       data,
	})
}

func ResponseData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       data,
	})
}
