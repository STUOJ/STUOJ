package runner

import (
	"STUOJ/internal/infrastructure/judge0"
	"STUOJ/internal/infrastructure/repository/entity"
	"strconv"

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
			Id:   int64(language.Id),
			Name: language.Name,
		})
	}
	return runnerLanguages, nil
}

func runnerSubmissionToJudgeSubmission(submission RunnerSubmission) []judge0.JudgeSubmission {
	var judgeSubmissions = make([]judge0.JudgeSubmission, len(submission.Testcase))
	for i, testCase := range submission.Testcase {
		judgeSubmissions[i] = judge0.JudgeSubmission{
			SourceCode:     submission.SourceCode,
			LanguageId:     uint64(submission.LanguageId),
			Stdin:          testCase.Input,
			ExpectedOutput: testCase.ExpectedOutput,
			CPUTimeLimit:   submission.CPUTimeLimit,
			MemoryLimit:    uint64(submission.MemoryLimit),
		}
	}
	return judgeSubmissions
}

func judgeResultToRunnerResult(judgeResult []judge0.JudgeResult) RunnerResult {
	var runnerResult RunnerResult
	for _, result := range judgeResult {
		status := RunnerStatus{
			Id:          int64(result.Status.Id),
			Description: result.Status.Description,
		}
		time, _ := strconv.ParseFloat(result.Time, 64)
		runnerResult.TestResult = append(runnerResult.TestResult, RunnerTestcaseResult{
			Stdout:  result.Stdout,
			Time:    time,
			Memory:  result.Memory,
			Stderr:  result.Stderr,
			Message: result.Message,
			Status:  status,
		})
		if result.Status.Id > uint64(runnerResult.Status.Id) {
			runnerResult.Status = status
		}
		if timeFloat, _ := strconv.ParseFloat(result.Time, 64); timeFloat > runnerResult.Time {
			runnerResult.Time = timeFloat
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
	if runnerResult.Status.Id < int64(entity.JudgeAC) {
		runnerResult.Status.Id = int64(entity.JudgeIE)
		runnerResult.Status.Description = entity.JudgeIE.String()
	}
	return runnerResult
}
