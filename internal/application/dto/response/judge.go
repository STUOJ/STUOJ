package response

import "STUOJ/internal/domain/runner"

type TestRunRes struct {
	Memory        int64   `json:"memory"`
	Status        int64   `json:"status"`
	Stdout        string  `json:"stdout"`
	Time          float64 `json:"time"`
	Stderr        string  `json:"stderr"`
	CompileOutput string  `json:"compile_output"`
}

func Result2TestRunRes(result runner.RunnerResult) TestRunRes {
	return TestRunRes{
		Memory:        int64(result.Memory),
		Status:        int64(result.Status.Id),
		Stdout:        result.TestResult[0].Stdout,
		Time:          result.Time,
		Stderr:        result.TestResult[0].Stderr,
		CompileOutput: result.CompileOutput,
	}
}
