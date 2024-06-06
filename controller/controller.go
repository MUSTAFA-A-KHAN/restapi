// controller/controller.go
package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mustafa-a-khan/restapi/models"
)

var Movies = []models.Movie{}

// Hello is a handler function that returns "Hello, World!" response.
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// StartApp is a handler function that returns "App Started" response.
func StartApp(c echo.Context) error {
	return c.String(http.StatusOK, "App Started")
}

func CreateMovie(c echo.Context) error {
	movie := new(models.Movie)
	c.Bind(movie)
	Movies = append(Movies, *movie)
	return c.JSON(200, Movies)
}

func ReadMovie(c echo.Context) error {
	return c.JSON(200, Movies)
}
func UpdateMovie(c echo.Context) error {
	movie := new(models.Movie)
	c.Bind(movie)
	for i := range Movies {
		if Movies[i].ID == movie.ID {
			Movies[i].Name = movie.Name
		}
	}
	return c.JSON(200, Movies)
}
func DeleteMovie(c echo.Context) error {
	movie := new(models.Movie)
	c.Bind(movie)
	for i := range Movies {
		if Movies[i].ID == movie.ID {
			Movies = append(Movies[:i], Movies[i+1:]...)
		}
	}
	return c.JSON(200, Movies)
}
