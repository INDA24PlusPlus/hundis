package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo) {
	/*
		These endpoints are used to solely compile code,
		the point is that these will be used to compile grading programs.

		TODO: Upload the compiled code to internal S3 bucket
	*/
	compile := e.Group("/compile")
	compile.POST("/cpp", CompileCpp)

	/*
		Endpoints to submit code for regular submissions
		Will synchronously compile the code then returd a job id for the running itself
	*/
	submit := e.Group("/submit")
	submit.POST("/cpp", SubmitCpp)
}
