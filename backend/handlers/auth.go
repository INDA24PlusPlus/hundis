package handlers

import (
	"hundis/config"
	"hundis/services/github"
	"hundis/services/identity"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GitHubAuth(c echo.Context) error {
	config := config.Config()

	code := c.QueryParam("code")
	next := c.QueryParam("state")

	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Code is required",
		})
	}

	tokenResponse, err := github.RequestToken(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get access token: " + err.Error(),
		})
	}

	userInfo, err := github.GetUserInfo(tokenResponse.AccessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get user info: " + err.Error(),
		})
	}

	user, err := identity.GetOrCreateUserByGitHubId(userInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get or create user: " + err.Error(),
		})
	}

	jwtToken, err := identity.CreateJWT(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to create JWT: " + err.Error(),
		})
	}
	c.SetCookie(&http.Cookie{
		Name:  "access_token",
		Value: jwtToken,
		Path:  "/",
	})

	return c.Redirect(http.StatusFound, config.BaseURL+next)
}
