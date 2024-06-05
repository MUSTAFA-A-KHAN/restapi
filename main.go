package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mustafa-a-khan/restapi/controller"
)

func main() {
	e := echo.New()

	e.GET("/hello", controller.Hello)
	e.GET("/", controller.StartApp)

	e.Start(":8080")
}
