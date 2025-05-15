package request

type QueryTestcaseParams struct {
	Page    *int64  `form:"page,omitempty"`
	Problem *string `form:"problem,omitempty"`
	Size    *int64  `form:"size,omitempty"`
}

type CreateTestcaseReq struct {
	ProblemId  int64  `json:"problem_id"`
	Serial     uint16 `json:"serial"`
	TestInput  string `json:"test_input"`
	TestOutput string `json:"test_output"`
}

type UpdateTestcaseReq struct {
	Id         int64  `json:"id"`
	ProblemId  int64  `json:"problem_id"`
	Serial     uint16 `json:"serial"`
	TestInput  string `json:"test_input"`
	TestOutput string `json:"test_output"`
}
