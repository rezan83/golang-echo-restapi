package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Album struct {
	ID     string
	Title  string
	Artist string
	Price  float32
}

var Albums = []Album{}

func homePageController(c echo.Context) error {
	return c.JSON(http.StatusOK, Albums)
}

func getAlbum(c echo.Context) error {
	id := c.Param("id")
	for _, album := range Albums {
		if album.ID == id {
			return c.JSON(http.StatusOK, album)

		}
	}
	return c.JSON(http.StatusNotFound, id)
}
func updateAlbum(c echo.Context) error {
	id := c.Param("id")
	newAlbum := new(Album)
	if err := c.Bind(newAlbum); err != nil {
		return err
	}
	for index, album := range Albums {
		if album.ID == id {
			Albums[index] = *newAlbum
			return c.JSON(http.StatusOK, newAlbum)

		}
	}
	return c.JSON(http.StatusNotFound, id)
}
func deleteAlbum(c echo.Context) error {
	id := c.Param("id")
	for index, album := range Albums {
		if album.ID == id {
			Albums = append(Albums[:index], Albums[index+1:]...)
			return c.JSON(http.StatusOK, Albums)

		}
	}
	return c.JSON(http.StatusNotFound, id)
}
func saveAlbum(c echo.Context) error {
	newAlbum := new(Album)
	if err := c.Bind(newAlbum); err != nil {
		return err
	}
	Albums = append(Albums, *newAlbum)
	return c.JSON(http.StatusOK, Albums)
}

func main() {
	fmt.Println("start")
	Albums = []Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	app := echo.New()

	app.GET("/", homePageController)
	app.POST("/albums", saveAlbum)
	app.GET("/albums/:id", getAlbum)
	app.PUT("/albums/:id", updateAlbum)
	app.DELETE("/albums/:id", deleteAlbum)

	app.Logger.Fatal(app.Start(":8585"))
}
