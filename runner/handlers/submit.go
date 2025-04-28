package handlers

import (
	"runner/compile"
	"runner/submit"

	"github.com/labstack/echo/v4"
)

type SubmitRequest struct {
	Code      string `json:"code"`
	ContestID string `json:"contest_id"`
	ProblemID string `json:"problem_id"`
}

func SubmitCpp(c echo.Context) error {
	var req SubmitRequest
	if err := c.Bind(&req); err != nil {
		return c.String(400, "Invalid request")
	}

	config, exists := compile.GetConfig("cpp")
	if !exists {
		return c.String(500, "Language not supported")
	}

	path, err := compile.Compile(req.Code, config)
	if err != nil {
		return c.String(500, "Compilation failed: "+err.Error())
	}

	id := submit.CreateSubmitJob(submit.Request{
		Path:      path,
		ContestID: req.ContestID,
		ProblemID: req.ProblemID,
		Config:    config,
	})

	return c.String(200, id)
}
