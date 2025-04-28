package dto

type Problem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Contest struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Problems    []Problem `json:"problems"`
}

type CreateContest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}
