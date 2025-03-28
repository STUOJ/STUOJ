package runner

import (
	judge0 "STUOJ/internal/domain/runner/judge0"
	"STUOJ/internal/domain/runner/valueobject"
)

type CodeRunner interface {
	CodeRun(valueobject.RunnerSubmission) valueobject.RunnerResult
	GetLanguage() ([]valueobject.RunnerLanguage, error)
}

var Runner CodeRunner = new(judge0.Judge0)
