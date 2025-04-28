package db

import (
	"database/sql"
	"hundis/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateRolesAndPermissions(db *gorm.DB) {
	// Define permissions
	var adminPerm = model.Permission{ID: 1, Name: "admin"}
	var createContestPerm = model.Permission{ID: 2, Name: "create:contest"}

	var permissions = []model.Permission{adminPerm}
	for _, permission := range permissions {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name"}),
		}).Create(&permission)
	}

	// Define roles
	var roles = []model.Role{
		{ID: model.AdminRoleID, Name: "admin", Permissions: []model.Permission{adminPerm, createContestPerm}},
		{ID: model.UserRoleID, Name: "user", Permissions: []model.Permission{createContestPerm}},
	}

	for _, role := range roles {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name"}),
		}).Create(&role)

		var dbRole model.Role
		db.Where("name = ?", role.Name).First(&dbRole)
		if len(role.Permissions) > 0 {
			db.Model(&dbRole).Association("Permissions").Clear()
			db.Model(&dbRole).Association("Permissions").Replace(role.Permissions)
		}
	}
}

func CreateOpenContest(db *gorm.DB) {
	var openContest = model.Contest{
		Name:        "Open",
		Slug:        "open",
		Description: "Open contest for all users! on skibidi bruh",
		StartTime:   sql.NullTime{},
		EndTime:     sql.NullTime{},
	}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "slug"}},
		DoUpdates: clause.AssignmentColumns([]string{"description", "start_time", "end_time", "name"}),
	}).Create(&openContest)

	var testProblem model.Problem
	err := db.First(&testProblem).Error
	if err != nil {
		testProblem = model.Problem{
			Name:        "Test Problem",
			Description: "This is a test problem",
			AdminID:     1,
		}
		db.Create(&testProblem)

		db.Model(&openContest).Association("Problems").Append(&testProblem)
	}
}
