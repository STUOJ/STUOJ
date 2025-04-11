package response

import "STUOJ/internal/domain/problem"

type ProblemData struct {
	CreateTime   string           `json:"create_time,omitempty"`
	Description  string           `json:"description,omitempty"`
	Difficulty   int64            `json:"difficulty"`
	Hint         string           `json:"hint,omitempty"`
	ID           int64            `json:"id"`
	Input        string           `json:"input,omitempty"`
	MemoryLimit  int64            `json:"memory_limit,omitempty"`
	Output       string           `json:"output,omitempty"`
	SampleInput  string           `json:"sample_input,omitempty"`
	SampleOutput string           `json:"sample_output,omitempty"`
	Source       string           `json:"source,omitempty"`
	Status       int64            `json:"status,omitempty"`
	TagIDS       []int64          `json:"tag_ids"`
	TimeLimit    int64            `json:"time_limit,omitempty"`
	Title        string           `json:"title"`
	UpdateTime   string           `json:"update_time,omitempty"`
	User         []UserSimpleData `json:"user"`
	ProblemUserScore
}

type ProblemSimpleData struct {
	Difficulty int64   `json:"difficulty,omitempty"`
	ID         int64   `json:"id"`
	Source     string  `json:"source,omitempty"`
	Title      string  `json:"title"`
	TagIDS     []int64 `json:"tag_ids"`
}

func Domain2ProblemSimpleData(p problem.Problem) ProblemSimpleData {
	return ProblemSimpleData{
		Difficulty: int64(p.Difficulty),
		ID:         int64(p.Id),
		Source:     p.Source.String(),
		Title:      p.Title.String(),
	}
}

type ProblemListItemData struct {
	CreateTime string  `json:"create_time,omitempty"`
	Difficulty int64   `json:"difficulty"`
	ID         int64   `json:"id"`
	Source     string  `json:"source,omitempty"`
	Status     int64   `json:"status,omitempty"`
	TagIDS     []int64 `json:"tag_ids"`
	Title      string  `json:"title"`
	UpdateTime string  `json:"update_time,omitempty"`
	ProblemUserScore
}

type ProblemUserScore struct {
	HasUserSubmission bool  `json:"has_user_submission"`
	UserScore         int64 `json:"user_score,omitempty"`
}
