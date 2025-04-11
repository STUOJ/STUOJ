package response

type TestcaseData struct {
	ID         *int64 `json:"id,omitempty"`
	ProblemID  int64  `json:"problem_id"`
	Serial     int64  `json:"serial"`
	TestInput  string `json:"test_input"`
	TestOutput string `json:"test_output"`
}
