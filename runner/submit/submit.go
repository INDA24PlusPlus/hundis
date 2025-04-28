package submit

import (
	"fmt"
	"runner/compile"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	Path      string          `json:"submission_path"`
	ContestID string          `json:"contest_id"`
	ProblemID string          `json:"problem_id"`
	Config    *compile.Config `json:"config"`
}

func CreateSubmitJob(request Request) string {
	id := uuid.New().String()
	go worker(request, id)
	return id
}

func worker(request Request, uuid string) {
	for range 10 {
		// Simulate some work
		time.Sleep(1 * time.Second)
		fmt.Println("Working on submission:", uuid)
	}
}
