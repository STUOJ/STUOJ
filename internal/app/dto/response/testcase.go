package response

type TestcaseData struct {
	Id         *int64 `json:"id,omitempty"`
	ProblemId  int64  `json:"problem_id"`
	Serial     int64  `json:"serial"`
	TestInput  string `json:"test_input"`
	TestOutput string `json:"test_output"`
}
