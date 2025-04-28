package handlers

import (
	"runner/compile"

	"github.com/labstack/echo/v4"
)

type CompileRequest struct {
	Code string `json:"code"`
}

func CompileCpp(c echo.Context) error {
	var req CompileRequest
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

	// TODO: Upload the file from path to S3 bucket and return some URI/URL

	return c.String(200, path)
}
