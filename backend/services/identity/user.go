package identity

import (
	"errors"
	"fmt"
	"hundis/db"
	"hundis/model"
	"hundis/services/github"

	"gorm.io/gorm"
)

func GetOrCreateUserByGitHubId(gitHubUserInfo *github.UserInfo) (model.User, error) {
	db := db.DB()
	var user model.User
	if err := db.Where("github_id = ?", gitHubUserInfo.ID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			username := GenerateUsername(gitHubUserInfo.Username)
			if username == "" {
				return model.User{}, errors.New("failed to generate username")
			}

			user = model.User{
				GitHubId:  gitHubUserInfo.ID,
				Username:  username,
				Email:     gitHubUserInfo.PrimaryEmail,
				AvatarURL: gitHubUserInfo.AvatarURL,
			}
			if err := db.Create(&user).Error; err != nil {
				return model.User{}, err
			}
		} else {
			return model.User{}, err
		}
	}

	return user, nil
}

func GetUserByUsername(username string) (model.User, error) {
	db := db.DB()
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func GenerateUsername(githubUsername string) string {
	db := db.DB()

	var username = githubUsername

	var userByUsername model.User
	if err := db.Where("username = ?", username).First(&userByUsername).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return githubUsername
		}
		return ""
	}

	for i := 1; ; i++ {
		candidateUsername := githubUsername + "_" + fmt.Sprintf("%d", i)

		if err := db.Where("username = ?", candidateUsername).First(&userByUsername).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return candidateUsername
			}
			return ""
		}
	}
}

func GetUserById(id uint) (model.User, error) {
	db := db.DB()
	var user model.User
	if err := db.Model(&model.User{}).Preload("Role.Permissions").Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func HasPermission(user model.User, permission string) bool {
	for _, perm := range user.Role.Permissions {
		if perm.Name == permission {
			return true
		}
	}
	return false
}

func UpdateAccount(userId uint, updateUser model.User) error {
	db := db.DB()

	var user model.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}

	user.Username = updateUser.Username
	user.Email = updateUser.Email

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
