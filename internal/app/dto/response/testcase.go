package response

type TestcaseData struct {
	Id         uint64 `json:"id,omitempty"`
	ProblemId  uint64 `json:"problem_id"`
	Serial     uint16 `json:"serial"`
	TestInput  string `json:"test_input"`
	TestOutput string `json:"test_output"`
}
