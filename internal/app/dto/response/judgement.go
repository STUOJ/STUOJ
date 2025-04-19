package response

type JudgementData struct {
	CompileOutput string  `json:"compile_output"`
	Id            int64   `json:"id"`
	Memory        int64   `json:"memory"`
	Message       string  `json:"message"`
	Status        int64   `json:"status"`
	Stderr        string  `json:"stderr"`
	Stdout        string  `json:"stdout"`
	SubmissionId  int64   `json:"submission_id"`
	TestcaseId    int64   `json:"testcase_id"`
	Time          float64 `json:"time"`
}
