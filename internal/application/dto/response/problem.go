package response

import (
	"STUOJ/internal/domain/problem"
	"STUOJ/pkg/utils"
)

type ProblemData struct {
	CreateTime   string `json:"create_time,omitempty"`
	Description  string `json:"description,omitempty"`
	Difficulty   int64  `json:"difficulty"`
	Hint         string `json:"hint,omitempty"`
	Id           int64  `json:"id"`
	Input        string `json:"input,omitempty"`
	MemoryLimit  int64  `json:"memory_limit,omitempty"`
	Output       string `json:"output,omitempty"`
	SampleInput  string `json:"sample_input,omitempty"`
	SampleOutput string `json:"sample_output,omitempty"`
	Source       string `json:"source,omitempty"`
	Status       int64  `json:"status,omitempty"`
	TimeLimit    int64  `json:"time_limit,omitempty"`
	Title        string `json:"title"`
	UpdateTime   string `json:"update_time,omitempty"`
	TagIds
}

type ProblemSimpleData struct {
	Difficulty int64  `json:"difficulty,omitempty"`
	Id         int64  `json:"id"`
	Source     string `json:"source,omitempty"`
	Title      string `json:"title"`
	TagIds
}

func Domain2ProblemSimpleData(p problem.Problem) ProblemSimpleData {
	return ProblemSimpleData{
		Difficulty: int64(p.Difficulty.Value()),
		Id:         int64(p.Id.Value()),
		Source:     p.Source.String(),
		Title:      p.Title.String(),
	}
}

func Map2ProblemSimpleData(p map[string]any) ProblemSimpleData {
	return ProblemSimpleData{
		Difficulty: p["difficulty"].(int64),
		Id:         p["id"].(int64),
		Source:     p["source"].(string),
		Title:      p["title"].(string),
	}
}

type ProblemQueryData struct {
	ProblemData
	ProblemUserScore
	User []UserSimpleData `json:"user"`
}

type ProblemListItemData struct {
	CreateTime string `json:"create_time,omitempty"`
	Difficulty int64  `json:"difficulty"`
	Id         int64  `json:"id"`
	Source     string `json:"source,omitempty"`
	Status     int64  `json:"status,omitempty"`
	Title      string `json:"title"`
	UpdateTime string `json:"update_time,omitempty"`
	ProblemUserScore
	TagIds
}

func Domain2ProblemListItemData(p problem.Problem) ProblemListItemData {
	return ProblemListItemData{
		CreateTime: p.CreateTime.String(),
		Difficulty: int64(p.Difficulty.Value()),
		Id:         int64(p.Id.Value()),
		Source:     p.Source.String(),
		Status:     int64(p.Status.Value()),
		Title:      p.Title.String(),
		UpdateTime: p.UpdateTime.String(),
	}
}

type ProblemUserScore struct {
	HasUserSubmission bool  `json:"has_user_submission"`
	UserScore         int64 `json:"user_score,omitempty"`
}

func Map2ProblemUserScore(p map[string]any) ProblemUserScore {
	var score ProblemUserScore
	score.HasUserSubmission = p["has_user_submission"].(int64) == 1
	if p["user_score"] != nil {
		score.UserScore = p["user_score"].(int64)
	}
	return score
}

type TagIds struct {
	TagIds []int64 `json:"tag_ids"`
}

// Map2TagIds 将map数据转换为TagIds结构体
// 当tag_ids为nil时返回空切片，避免panic
func Map2TagIds(p map[string]any) TagIds {
	if p["problem_tag_id"] == nil {
		return TagIds{TagIds: []int64{}}
	}
	tagIds, _ := utils.StringToInt64Slice(p["problem_tag_id"].(string))
	return TagIds{
		TagIds: tagIds,
	}
}
