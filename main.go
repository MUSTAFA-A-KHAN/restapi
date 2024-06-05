package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mustafa-a-khan/restapi/controller"
)

func main() {
	e := echo.New()

	e.GET("/hello", controller.Hello)
	e.GET("/", controller.StartApp)
	e.POST("/Movie", controller.CreateMovie)
	e.PUT("/Movie", controller.UpdateMovie)
	e.GET("/Movie", controller.ReadMovie)
	e.DELETE("/Movie", controller.DeleteMovie)
	e.Start(":8080")
}
