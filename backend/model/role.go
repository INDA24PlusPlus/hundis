package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint         `gorm:"primarykey"`
	Name        string       `json:"name" gorm:"unique"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
	Users       []User       `json:"users"`
}

type Permission struct {
	gorm.Model
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name" gorm:"unique"`
}
