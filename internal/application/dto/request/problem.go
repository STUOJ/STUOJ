package request

type QueryProblemParams struct {
	Difficulty *string `form:"difficulty,omitempty"`
	EndTime    *string `form:"end-time,omitempty"`
	Order      *string `form:"order,omitempty"`
	OrderBy    *string `form:"order_by,omitempty"`
	Page       *int64  `form:"page,omitempty"`
	Size       *int64  `form:"size,omitempty"`
	StartTime  *string `form:"start-time,omitempty"`
	Status     *string `form:"status,omitempty"`
	Tag        *string `form:"tag,omitempty"`
	Title      *string `form:"title,omitempty"`
	User       *int64  `form:"user,omitempty"`
}

type ProblemStatisticsParams struct {
	QueryProblemParams
	GroupBy string `form:"group_by"`
}

type QueryOneProblemParams struct {
	Solutions *bool `form:"solutions,omitempty"`
	Testcases *bool `form:"testcases,omitempty"`
}

type CreateProblemReq struct {
	Description  string  `json:"description,omitempty"`
	Difficulty   int64   `json:"difficulty"`
	Hint         string  `json:"hint,omitempty"`
	Input        string  `json:"input,omitempty"`
	MemoryLimit  int64   `json:"memory_limit,omitempty"`
	Output       string  `json:"output,omitempty"`
	SampleInput  string  `json:"sample_input,omitempty"`
	SampleOutput string  `json:"sample_output,omitempty"`
	Source       string  `json:"source,omitempty"`
	Status       int64   `json:"status,omitempty"`
	TagIds       []int64 `json:"tag_ids"`
	TimeLimit    int64   `json:"time_limit,omitempty"`
	Title        string  `json:"title"`
}

type UpdateProblemReq struct {
	Description  string  `json:"description,omitempty"`
	Difficulty   int64   `json:"difficulty"`
	Hint         string  `json:"hint,omitempty"`
	Id           int64   `json:"id"`
	Input        string  `json:"input,omitempty"`
	MemoryLimit  int64   `json:"memory_limit,omitempty"`
	Output       string  `json:"output,omitempty"`
	SampleInput  string  `json:"sample_input,omitempty"`
	SampleOutput string  `json:"sample_output,omitempty"`
	Source       string  `json:"source,omitempty"`
	Status       int64   `json:"status,omitempty"`
	TagIds       []int64 `json:"tag_ids"`
	TimeLimit    int64   `json:"time_limit,omitempty"`
	Title        string  `json:"title"`
}
