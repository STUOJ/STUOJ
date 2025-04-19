package runner

type RunnerSubmission struct {
	SourceCode   string                     `json:"source_code"`
	LanguageId   int64                      `json:"language_id"`
	Testcase     []RunnerTestcaseSubmission `json:"testcase,omitempty"`
	CPUTimeLimit float64                    `json:"cpu_time_limit,omitempty"`
	MemoryLimit  int64                      `json:"memory_limit,omitempty"`
}

type RunnerTestcaseSubmission struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output,omitempty"`
}

type RunnerResult struct {
	Status        RunnerStatus           `json:"status"`
	TestResult    []RunnerTestcaseResult `json:"test_result"`
	CompileOutput string                 `json:"compile_output,omitempty"`
	Time          float64                `json:"time,omitempty"`
	Memory        float64                `json:"memory,omitempty"`
	Message       string                 `json:"message,omitempty"`
}

type RunnerTestcaseResult struct {
	Stdout  string       `json:"stdout,omitempty"`
	Time    float64      `json:"time,omitempty"`
	Memory  float64      `json:"memory,omitempty"`
	Stderr  string       `json:"stderr,omitempty"`
	Message string       `json:"message,omitempty"`
	Status  RunnerStatus `json:"status"`
}

type RunnerStatus struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
}

type RunnerLanguage struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CodeRunner interface {
	CodeRun(RunnerSubmission) RunnerResult
	GetLanguage() ([]RunnerLanguage, error)
}

var Runner CodeRunner = new(Judge0)
