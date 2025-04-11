package response

type FpsProblem struct {
	CreateTime   string `json:"create_time"`
	Description  string `json:"description"`
	Input        string `json:"input"`
	MemoryLimit  int64  `json:"memory_limit"`
	Output       string `json:"output"`
	SampleInput  string `json:"sample_input"`
	SampleOutput string `json:"sample_output"`
	TimeLimit    int64  `json:"time_limit"`
	Title        string `json:"title"`
	UpdateTime   string `json:"update_time"`
}

// fps题解
type FpsSolution struct {
	CreateTime string `json:"create_time"`
	LanguageID int64  `json:"language_id"`
	SourceCode string `json:"source_code"`
	UpdateTime string `json:"update_time"`
}

// fps测试数据
type FpsTestcase struct {
	TestInput  *string `json:"test_input,omitempty"`
	TestOutput *string `json:"test_output,omitempty"`
}
