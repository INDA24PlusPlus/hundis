package handlers

import (
	"errors"
	"hundis/model"
	"hundis/services/identity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Me(c echo.Context) error {
	userId := c.Get("userId").(uint)

	user, err := identity.GetUserById(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to get user: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":          user.ID,
		"username":    user.Username,
		"email":       user.Email,
		"avatarUrl":   user.AvatarURL,
		"permissions": mapPermissionNames(user.Role.Permissions),
	})
}

func mapPermissionNames(permissions []model.Permission) []string {
	names := make([]string, len(permissions))
	for i, perm := range permissions {
		names[i] = perm.Name
	}
	return names
}

func UpdateAccount(c echo.Context) error {
	userId := c.Get("userId").(uint)

	var updateUser model.User
	if err := c.Bind(&updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid input: " + err.Error(),
		})
	}

	if len(updateUser.Username) < 4 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Username must be at least 4 characters long",
		})
	}
	if len(updateUser.Username) > 30 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Username must be at most 30 characters long",
		})
	}
	if len(updateUser.Email) < 3 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Email must be at least 3 characters long",
		})
	}
	if len(updateUser.Email) > 100 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Email must be at most 100 characters long",
		})
	}

	err := identity.UpdateAccount(userId, updateUser)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Username already exists",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Failed to update profile",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Profile updated successfully",
	})
}
