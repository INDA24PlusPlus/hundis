package middlewares

import (
	"hundis/services/identity"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorized := checkAuth(c)
		if !authorized {
			return c.String(401, "Unauthorized")
		}

		return next(c)
	}
}

func AuthWithPermissions(permissions []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorized := checkAuth(c)
			if !authorized {
				return c.String(401, "Unauthorized")
			}

			userId := c.Get("userId").(uint)
			user, err := identity.GetUserById(userId)

			if err != nil {
				return c.String(500, "Internal Server Error")
			}

			for _, permission := range permissions {
				if !identity.HasPermission(user, permission) {
					return c.String(403, "Forbidden")
				}
			}

			return next(c)
		}
	}
}

func checkAuth(c echo.Context) bool {
	accessToken := c.Request().Header.Get("Authorization")
	if accessToken == "" {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			return false
		}
		accessToken = cookie.Value
	} else {
		if len(accessToken) < 7 {
			return false
		}
		accessToken = accessToken[7:]
	}

	userID, err := identity.ValidateJWT(accessToken)
	if err != nil {
		return false
	}
	c.Set("userId", userID)

	return true
}
