package response

import "STUOJ/internal/domain/history"

type HistoryData struct {
	CreateTime   string         `json:"create_time"`
	Description  string         `json:"description"`
	Difficulty   int64          `json:"difficulty"`
	Hint         string         `json:"hint"`
	Id           int64          `json:"id"`
	Input        string         `json:"input"`
	MemoryLimit  int64          `json:"memory_limit"`
	Operation    int64          `json:"operation"`
	Output       string         `json:"output"`
	ProblemId    int64          `json:"problem_id"`
	SampleInput  string         `json:"sample_input"`
	SampleOutput string         `json:"sample_output"`
	Source       string         `json:"source"`
	TimeLimit    int64          `json:"time_limit"`
	Title        string         `json:"title"`
	User         UserSimpleData `json:"user"`
}

func Domain2HistoryData(h history.History) HistoryData {
	return HistoryData{
		Id:           h.Id.Value(),
		ProblemId:    h.ProblemId.Value(),
		Title:        h.Title.String(),
		Source:       h.Source.String(),
		Difficulty:   int64(h.Difficulty.Value()),
		TimeLimit:    int64(h.TimeLimit.Value()),
		MemoryLimit:  h.MemoryLimit.Value(),
		Description:  h.Description.String(),
		Input:        h.Input.String(),
		Output:       h.Output.String(),
		SampleInput:  h.SampleInput.String(),
		SampleOutput: h.SampleOutput.String(),
		Hint:         h.Hint.String(),
		Operation:    int64(h.Operation),
		CreateTime:   h.CreateTime.String(),
	}
}

type HistoryListItemData struct {
	CreateTime string            `json:"create_time"`
	Difficulty int64             `json:"difficulty"`
	Id         int64             `json:"id"`
	Operation  int64             `json:"operation"`
	ProblemID  int64             `json:"problem_id"`
	Title      string            `json:"title"`
	User       UserSimpleData    `json:"user"`
	Problem    ProblemSimpleData `json:"problem"`
}

func Domain2HistoryListItem(h history.History) HistoryListItemData {
	return HistoryListItemData{
		Id:         h.Id.Value(),
		ProblemID:  h.ProblemId.Value(),
		Title:      h.Title.String(),
		Difficulty: int64(h.Difficulty.Value()),
		Operation:  int64(h.Operation),
		CreateTime: h.CreateTime.String(),
	}
}
