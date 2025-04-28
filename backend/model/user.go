package model

import (
	"gorm.io/gorm"
)

const (
	AdminRoleID = 1
	UserRoleID  = 2
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"uniqueIndex:idx_users_lower_username,expression:lower(username)"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarUrl"`
	GitHubId  int    `json:"githubId" gorm:"unique;column:github_id"`
	RoleID    uint   `json:"roleId" gorm:"default:2"` // default role is user
	Role      Role   `json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.RoleID == 0 {
		u.RoleID = UserRoleID
	}
	return nil
}
