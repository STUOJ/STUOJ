package request

type CreateTestcaseReq struct {
	ProblemId  uint64 `json:"problem_id"`
	Serial     uint16 `json:"serial"`
	TestInput  string `json:"test_input"`
	TestOutput string `json:"test_output"`
}

type UpdateTestcaseReq struct {
	Id         uint64 `json:"id"`
	Serial     uint16 `json:"serial"`
	TestInput  string `json:"test_input"`
	TestOutput string `json:"test_output"`
}
