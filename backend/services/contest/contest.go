package contest

import (
	"hundis/db"
	"hundis/dto"
	"hundis/model"
)

func GetContestById(id string) (model.Contest, error) {
	db := db.DB()
	var contest model.Contest
	if err := db.Where("id = ?", id).Preload("Problems").First(&contest).Error; err != nil {
		return model.Contest{}, err
	}

	return contest, nil
}

func GetContests() ([]model.Contest, error) {
	db := db.DB()
	var contests []model.Contest
	if err := db.Order("id").Find(&contests).Error; err != nil {
		return nil, err
	}

	return contests, nil
}

func CreateContest(contest dto.CreateContest) (model.Contest, error) {
	db := db.DB()
	newContest := model.Contest{
		Name:        contest.Name,
		Slug:        contest.Slug,
		Description: contest.Description,
	}

	if err := db.Create(&newContest).Error; err != nil {
		return model.Contest{}, err
	}

	return newContest, nil
}

func UpdateContest(id string, contest dto.CreateContest) (model.Contest, error) {
	db := db.DB()
	existingContest, err := GetContestById(id)
	if err != nil {
		return model.Contest{}, err
	}

	existingContest.Name = contest.Name
	existingContest.Slug = contest.Slug
	existingContest.Description = contest.Description

	if err := db.Save(&existingContest).Error; err != nil {
		return model.Contest{}, err
	}

	return existingContest, nil
}
