package handlers

import (
	"errors"
	"hundis/dto"
	"hundis/services/contest"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetContest(c echo.Context) error {
	contestId := c.Param("id")
	contest, err := contest.GetContestById(contestId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(404, map[string]interface{}{
				"error": "Contest not found",
			})
		}

		return c.JSON(500, map[string]interface{}{
			"error": "Failed to get contest: " + err.Error(),
		})
	}

	var problems []dto.Problem = []dto.Problem{}
	for _, problem := range contest.Problems {
		problems = append(problems, dto.Problem{
			ID:   problem.ID,
			Name: problem.Name,
		})
	}

	response := dto.Contest{
		ID:          contest.ID,
		Name:        contest.Name,
		Slug:        contest.Slug,
		Description: contest.Description,
		Problems:    problems,
	}

	return c.JSON(200, response)
}

func GetContests(c echo.Context) error {
	contests, err := contest.GetContests()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": "Failed to get contests: " + err.Error(),
		})
	}

	var contestsDTO []dto.Contest
	for _, contest := range contests {
		contestsDTO = append(contestsDTO, dto.Contest{
			ID:          contest.ID,
			Name:        contest.Name,
			Slug:        contest.Slug,
			Description: contest.Description,
			Problems:    []dto.Problem{},
		})
	}

	return c.JSON(200, contestsDTO)
}

func CreateContest(c echo.Context) error {
	var createContestDTO dto.CreateContest
	if err := c.Bind(&createContestDTO); err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	if len(createContestDTO.Name) < 5 || len(createContestDTO.Name) > 100 {
		return c.JSON(400, map[string]interface{}{
			"error": "Contest name must be between 5 and 100 characters long",
		})
	}

	if len(createContestDTO.Slug) < 5 || len(createContestDTO.Slug) > 30 {
		return c.JSON(400, map[string]interface{}{
			"error": "Contest slug must be between 5 and 30 characters long",
		})
	}

	contest, err := contest.CreateContest(createContestDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(400, map[string]interface{}{
				"error": "Contest with this slug already exists",
			})
		}

		return c.JSON(500, map[string]interface{}{
			"error": "Failed to create contest: " + err.Error(),
		})
	}

	contestDTO := dto.Contest{
		ID:          contest.ID,
		Name:        contest.Name,
		Slug:        contest.Slug,
		Description: contest.Description,
		Problems:    []dto.Problem{},
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Contest created successfully",
		"contest": contestDTO,
	})
}

func UpdateContest(c echo.Context) error {
	contestId := c.Param("id")
	var updateContestDTO dto.CreateContest
	if err := c.Bind(&updateContestDTO); err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	if len(updateContestDTO.Name) < 5 || len(updateContestDTO.Name) > 100 {
		return c.JSON(400, map[string]interface{}{
			"error": "Contest name must be between 5 and 100 characters long",
		})
	}

	if len(updateContestDTO.Slug) < 5 || len(updateContestDTO.Slug) > 30 {
		return c.JSON(400, map[string]interface{}{
			"error": "Contest slug must be between 5 and 30 characters long",
		})
	}

	contest, err := contest.UpdateContest(contestId, updateContestDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(404, map[string]interface{}{
				"error": "Contest not found",
			})
		}

		return c.JSON(500, map[string]interface{}{
			"error": "Failed to update contest: " + err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Contest updated successfully",
		"contest": dto.Contest{
			ID:          contest.ID,
			Name:        contest.Name,
			Slug:        contest.Slug,
			Description: contest.Description,
			Problems:    []dto.Problem{},
		},
	})
}
