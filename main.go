package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mustafa-a-khan/restapi/controller"
	"github.com/mustafa-a-khan/restapi/models"
	"github.com/mustafa-a-khan/restapi/services"
)

// https://gogoanime3.co/category/one-piece
func main() {
	fmt.Print("inside main")

	ScrapeReq := models.ScrapeRequest{
		URL: "gogoanime3.co",
		Parameters: map[string]string{
			"path1": "one-piece-episode-99",
			"path2": "",
		},
	}

	fmt.Print("AFTER MAIN")

	services.Test(ScrapeReq)

	e := echo.New()

	e.GET("/hello", controller.Hello)
	e.GET("/", controller.StartApp)
	e.POST("/Movie/details", controller.Callmovie)
	e.POST("/Movie", controller.CreateMovie)
	e.PUT("/Movie", controller.UpdateMovie)
	e.GET("/Movie", controller.ReadMovie)
	e.DELETE("/Movie", controller.DeleteMovie)

	e.Start(":8080")
}
