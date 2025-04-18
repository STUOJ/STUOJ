package response

import (
	"STUOJ/internal/domain/problem"
	"STUOJ/utils"
)

type ProblemData struct {
	CreateTime   string           `json:"create_time,omitempty"`
	Description  string           `json:"description,omitempty"`
	Difficulty   int64            `json:"difficulty"`
	Hint         string           `json:"hint,omitempty"`
	Id           int64            `json:"id"`
	Input        string           `json:"input,omitempty"`
	MemoryLimit  int64            `json:"memory_limit,omitempty"`
	Output       string           `json:"output,omitempty"`
	SampleInput  string           `json:"sample_input,omitempty"`
	SampleOutput string           `json:"sample_output,omitempty"`
	Source       string           `json:"source,omitempty"`
	Status       int64            `json:"status,omitempty"`
	TagIdS       []int64          `json:"tag_ids"`
	TimeLimit    int64            `json:"time_limit,omitempty"`
	Title        string           `json:"title"`
	UpdateTime   string           `json:"update_time,omitempty"`
	User         []UserSimpleData `json:"user"`
}

type ProblemSimpleData struct {
	Difficulty int64   `json:"difficulty,omitempty"`
	Id         int64   `json:"id"`
	Source     string  `json:"source,omitempty"`
	Title      string  `json:"title"`
	TagIdS     []int64 `json:"tag_ids"`
}

func Domain2ProblemSimpleData(p problem.Problem) ProblemSimpleData {
	return ProblemSimpleData{
		Difficulty: int64(p.Difficulty),
		Id:         int64(p.Id),
		Source:     p.Source.String(),
		Title:      p.Title.String(),
	}
}

func Map2ProblemSimpleData(p map[string]any) ProblemSimpleData {
	tagIds, _ := utils.StringToInt64Slice(p["tag_ids"].(string))
	return ProblemSimpleData{
		Difficulty: p["difficulty"].(int64),
		Id:         p["id"].(int64),
		Source:     p["source"].(string),
		Title:      p["title"].(string),
		TagIdS:     tagIds,
	}
}

type ProblemQueryData struct {
	ProblemData
	ProblemUserScore
}

type ProblemListItemData struct {
	CreateTime string  `json:"create_time,omitempty"`
	Difficulty int64   `json:"difficulty"`
	Id         int64   `json:"id"`
	Source     string  `json:"source,omitempty"`
	Status     int64   `json:"status,omitempty"`
	TagIdS     []int64 `json:"tag_ids"`
	Title      string  `json:"title"`
	UpdateTime string  `json:"update_time,omitempty"`
	ProblemUserScore
}

type ProblemUserScore struct {
	HasUserSubmission bool  `json:"has_user_submission"`
	UserScore         int64 `json:"user_score,omitempty"`
}

func Map2ProblemUserScore(p map[string]any) ProblemUserScore {
	var score ProblemUserScore
	score.HasUserSubmission = p["has_user_submission"].(bool)
	if p["user_score"] != nil {
		score.UserScore = p["user_score"].(int64)
	}
	return score
}
