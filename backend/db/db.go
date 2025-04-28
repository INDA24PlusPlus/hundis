package db

import (
	"fmt"
	"hundis/config"
	"hundis/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	var cfg = config.Config()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Stockholm", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("failed to connect database")
	}

	err := db.AutoMigrate(&model.Role{}, &model.Permission{})
	if err != nil {
		panic("Failed to migrate role and permission tables: " + err.Error())
	}

	CreateRolesAndPermissions(db)

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("Failed to migrate permission and role tables: " + err.Error())
	}

	err = db.AutoMigrate(&model.Contest{}, &model.Problem{}, &model.ContestUser{})
	if err != nil {
		panic("Failed to migrate contest and problem tables: " + err.Error())
	}

	CreateOpenContest(db)
}

func DB() *gorm.DB {
	return db
}
