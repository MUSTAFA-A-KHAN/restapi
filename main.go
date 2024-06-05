package main

import "github.com/labstack/echo/v4"

func hello(c echo.Context) error {
	return c.JSON(200, "Hello World")
}
func statrApp(c echo.Context) error {
	return c.JSON(200, "App has Started")
}

func main() {
	e := echo.New()

	e.GET("/hello", hello)
	e.GET("/hello", statrApp)

	e.Start(":8080")
}
