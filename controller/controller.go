// controller/controller.go
package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello is a handler function that returns "Hello, World!" response.
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// StartApp is a handler function that returns "App Started" response.
func StartApp(c echo.Context) error {
	return c.String(http.StatusOK, "App Started")
}
