package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	Name        string        `json:"name"`
	Slug        string        `json:"slug" gorm:"unique"`
	Description string        `json:"description"`
	StartTime   sql.NullTime  `json:"startTime"`
	EndTime     sql.NullTime  `json:"endTime"`
	Problems    []Problem     `json:"problems" gorm:"many2many:contest_problems"`
	Users       []ContestUser `json:"users" gorm:"many2many:contest_users"`
}

type ContestUser struct {
	gorm.Model
	Contest   Contest `json:"contest"`
	ContestID uint    `json:"contestId"`
	User      User    `json:"user"`
	UserID    uint    `json:"userId"`
	Role      Role    `json:"role"`
	RoleID    uint    `json:"roleId"`
}
