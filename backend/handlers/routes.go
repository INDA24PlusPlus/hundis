package handlers

import (
	"hundis/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.GET("/github", GitHubAuth)

	e.GET("/me", Me, middlewares.Auth)

	settings := e.Group("/settings")
	settings.PUT("/account", UpdateAccount, middlewares.Auth)

	contest := e.Group("/contests")
	contest.GET("", GetContests)
	contest.GET("/:id", GetContest)
	contest.POST("", CreateContest, middlewares.AuthWithPermissions([]string{"create:contest"}))
	contest.PUT("/:id", UpdateContest, middlewares.AuthWithPermissions([]string{"create:contest"}))
}
