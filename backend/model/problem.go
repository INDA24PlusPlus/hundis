package model

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Admin       User      `json:"admin"`
	AdminID     uint      `json:"adminId"`
	Description string    `json:"description"`
	Contests    []Contest `json:"contests" gorm:"many2many:contest_problems"`
}
