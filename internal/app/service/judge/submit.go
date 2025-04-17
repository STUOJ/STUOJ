package judge

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/judgement"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/runner"
	"STUOJ/internal/domain/submission"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"STUOJ/internal/model/querycontext"
	"STUOJ/utils"
	"slices"
)

func Submit(req request.JudgeReq, reqUser model.ReqUser) (uint64, error) {
	languageQuery := querycontext.LanguageQueryContext{}
	languageQuery.Id.Add(req.LanguageID)
	languageQuery.Field = *query.LanguageAllField
	languageDomain, _, err := language.Query.SelectOne(languageQuery)
	if err != nil {
		return 0, err
	}
	languageMapId := languageDomain.MapId

	problemQuery := querycontext.ProblemQueryContext{}
	problemQuery.Id.Add(req.ProblemID)
	problemDomain, problemMap, err := problem.Query.SelectOne(problemQuery)
	if err != nil {
		return 0, errors.ErrNotFound.WithMessage("题目不存在")
	}
	if problemDomain.Status < entity.ProblemPublic && reqUser.Role < entity.RoleAdmin {
		userIds, err := utils.StringToInt64Slice(problemMap["problem_user_id"].(string))
		if err != nil {
			return 0, errors.ErrInternalServer.WithMessage("内部错误")
		}
		if !slices.Contains(userIds, reqUser.ID) {
			return 0, errors.ErrNotFound.WithMessage("题目不存在")
		}
	}

	// 先创建提交记录，状态为等待评测，确保CreateTime是提交时间
	submissionDomain := submission.NewSubmission(
		submission.WithUserId(uint64(reqUser.ID)),
		submission.WithProblemId(uint64(req.ProblemID)),
		submission.WithStatus(entity.JudgeIE), // 设置为IE状态,保证在错误导致中断的情况下显示为IE
		submission.WithLength(uint32(len(req.SourceCode))),
		submission.WithLanguageId(uint64(languageDomain.Id)),
		submission.WithSourceCode(req.SourceCode),
	)
	submissionId, err := submissionDomain.Create()
	if err != nil {
		return 0, err
	}

	testcaseQuery := querycontext.TestcaseQueryContext{}
	testcaseQuery.ProblemId.Add(req.ProblemID)
	testcaseQuery.Page = option.NewPagination(0, 1000)
	testcaseQuery.Field = *query.TestcaseAllField
	testcaseDomain, _, err := testcase.Query.Select(testcaseQuery)
	if err != nil {
		return 0, errors.ErrNotFound.WithMessage("找不到测试点")
	}
	testcaseSubmission := make([]runner.RunnerTestcaseSubmission, len(testcaseDomain))
	for i, v := range testcaseDomain {
		testcaseSubmission[i] = runner.RunnerTestcaseSubmission{
			Input:          v.TestInput.String(),
			ExpectedOutput: v.TestOutput.String(),
		}
	}

	// 先创建空的judgement记录，确保CreateTime是提交时间
	judgementIds := make([]uint64, len(testcaseDomain))
	for i, tc := range testcaseDomain {
		judgementDomain := judgement.NewJudgement(
			judgement.WithSubmissionId(submissionId),
			judgement.WithTestcaseId(tc.Id),
			judgement.WithStatus(entity.JudgeIE), //设置为IE状态,保证在错误导致中断的情况下显示为IE
			judgement.WithCompileOutput(""),
			judgement.WithMessage(""),
		)
		judgementId, err := judgementDomain.Create()
		if err != nil {
			return 0, err
		}
		judgementIds[i] = judgementId
	}

	runnerSubmission := runner.RunnerSubmission{
		LanguageId:   uint64(languageMapId),
		SourceCode:   req.SourceCode,
		Testcase:     testcaseSubmission,
		CPUTimeLimit: problemDomain.TimeLimit,
		MemoryLimit:  problemDomain.MemoryLimit,
	}
	runnerResult := runner.Runner.CodeRun(runnerSubmission)

	var score uint8 = 0
	for i, v := range runnerResult.TestResult {
		if v.Status.Id == uint64(entity.JudgeAC) {
			score += uint8(100 / uint8(len(runnerResult.TestResult)))
		}
		// 更新judgement记录
		judgementDomain := judgement.NewJudgement(
			judgement.WithId(judgementIds[i]),
			judgement.WithTime(v.Time),
			judgement.WithMemory(uint64(v.Memory)),
			judgement.WithStdout(v.Stdout),
			judgement.WithStderr(v.Stderr),
			judgement.WithCompileOutput(runnerResult.CompileOutput),
			judgement.WithMessage(v.Message),
			judgement.WithStatus(entity.JudgeStatus(v.Status.Id)),
		)
		err = judgementDomain.Update()
		if err != nil {
			return 0, err
		}
	}

	// 更新提交记录的状态和评测结果
	submissionDomain = submission.NewSubmission(
		submission.WithId(submissionId),
		submission.WithStatus(entity.JudgeStatus(runnerResult.Status.Id)),
		submission.WithScore(score),
		submission.WithMemory(uint64(runnerResult.Memory)),
		submission.WithTime(runnerResult.Time),
	)
	err = submissionDomain.Update()
	if err != nil {
		return 0, err
	}
	return submissionId, nil
}
