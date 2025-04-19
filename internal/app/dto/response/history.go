package response

type HistoryData struct {
	CreateTime   string            `json:"create_time"`
	Description  string            `json:"description"`
	Difficulty   int64             `json:"difficulty"`
	Hint         string            `json:"hint"`
	Id           int64             `json:"id"`
	Input        string            `json:"input"`
	MemoryLimit  int64             `json:"memory_limit"`
	Operation    int64             `json:"operation"`
	Output       string            `json:"output"`
	ProblemId    int64             `json:"problem_id"`
	SampleInput  string            `json:"sample_input"`
	SampleOutput string            `json:"sample_output"`
	Source       string            `json:"source"`
	TimeLimit    int64             `json:"time_limit"`
	Title        string            `json:"title"`
	UserId       int64             `json:"user_id"`
	Problem      ProblemSimpleData `json:"problem"`
	User         UserSimpleData    `json:"user"`
}
