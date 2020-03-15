package impl

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Good!")
}

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "Welcome to his Service",
	})
}
