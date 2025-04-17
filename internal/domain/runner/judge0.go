package runner

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/db/entity"

	"sync"

	"golang.org/x/sync/errgroup"
)

type Judge0 struct{}

func (j *Judge0) CodeRun(runnerSubmission RunnerSubmission) RunnerResult {
	judgeSubmissions := runnerSubmissionToJudgeSubmission(runnerSubmission)
	runNum := len(judgeSubmissions)
	chJudgement := make(chan judge0.JudgeResult, runNum)

	var eg errgroup.Group
	var mu sync.Mutex // 用于保护错误信息的互斥锁
	var errs []error

	for i := range judgeSubmissions {
		submission := judgeSubmissions[i] // 避免闭包陷阱
		eg.Go(func() error {
			result, err := judge0.Submit(submission)
			if err != nil {
				result.Status.Id = uint64(entity.JudgeIE)
				result.Status.Description = entity.JudgeIE.String()
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			}
			chJudgement <- result
			return nil
		})
	}

	// 等待所有任务完成
	_ = eg.Wait()
	close(chJudgement)

	// 收集结果
	judgeResult := make([]judge0.JudgeResult, 0, runNum)
	for result := range chJudgement {
		judgeResult = append(judgeResult, result)
	}

	return judgeResultToRunnerResult(judgeResult)
}

func (j *Judge0) GetLanguage() ([]RunnerLanguage, error) {
	languages, err := judge0.GetLanguage()
	if err != nil {
		return nil, err
	}
	var runnerLanguages []RunnerLanguage
	for _, language := range languages {
		runnerLanguages = append(runnerLanguages, RunnerLanguage{
			Id:   language.Id,
			Name: language.Name,
		})
	}
	return runnerLanguages, nil
}

func runnerSubmissionToJudgeSubmission(submission RunnerSubmission) []judge0.JudgeSubmission {
	var judgeSubmissions = make([]judge0.JudgeSubmission, len(submission.Testcase))
	for _, testCase := range submission.Testcase {
		judgeSubmissions = append(judgeSubmissions, judge0.JudgeSubmission{
			SourceCode:     submission.SourceCode,
			LanguageId:     submission.LanguageId,
			Stdin:          testCase.Input,
			ExpectedOutput: testCase.ExpectedOutput,
			CPUTimeLimit:   submission.CPUTimeLimit,
			MemoryLimit:    submission.MemoryLimit,
		})
	}
	return judgeSubmissions
}

func judgeResultToRunnerResult(judgeResult []judge0.JudgeResult) RunnerResult {
	var runnerResult RunnerResult
	for _, result := range judgeResult {
		status := RunnerStatus{
			Id:          result.Status.Id,
			Description: result.Status.Description,
		}
		runnerResult.TestResult = append(runnerResult.TestResult, RunnerTestcaseResult{
			Stdout:  result.Stdout,
			Time:    result.Time,
			Memory:  result.Memory,
			Stderr:  result.Stderr,
			Message: result.Message,
			Status:  status,
		})
		if result.Status.Id > runnerResult.Status.Id {
			runnerResult.Status = status
		}
		if result.Time > runnerResult.Time {
			runnerResult.Time = result.Time
		}
		if result.Memory > runnerResult.Memory {
			runnerResult.Memory = result.Memory
		}
		if result.Message != "" {
			runnerResult.Message = result.Message
		}
		if result.CompileOutput != "" {
			runnerResult.CompileOutput = result.CompileOutput
		}
	}
	if runnerResult.Status.Id < uint64(entity.JudgeAC) {
		runnerResult.Status.Id = uint64(entity.JudgeIE)
		runnerResult.Status.Description = entity.JudgeIE.String()
	}
	return runnerResult
}
