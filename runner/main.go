package main

import (
	"runner/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	/* e.GET("/", func(c echo.Context) error {
		command := c.QueryParam("command")

		exec, err := exec.Command("nsjail", "-Mo", "--chroot", "/chroot/", "--", command).CombinedOutput()
		if err != nil {
			return c.String(500, string(exec))
		}
		return c.String(200, string(exec))
	}) */

	handlers.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":5001"))
}
