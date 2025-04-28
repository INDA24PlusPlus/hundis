package main

import (
	"hundis/config"
	"hundis/db"
	"hundis/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handlers.SetupRoutes(e)

	config.Init()
	db.Init()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello, World!",
		})
	})

	e.Logger.Fatal(e.Start(":5000"))
}
