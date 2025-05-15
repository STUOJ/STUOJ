package record

import (
	model "STUOJ/internal/application/dto"
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/judgement"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/submission"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/infrastructure/persistence/entity"
	querycontext2 "STUOJ/internal/infrastructure/persistence/repository/querycontext"
	query "STUOJ/internal/infrastructure/persistence/repository/queryfield"
)

type SubmissionPage struct {
	Submissions []response.SubmissionData `json:"submissions"`
	model.Page
}

// Select 查询所有提交记录
func Select(params request.QuerySubmissionParams, reqUser request.ReqUser) (SubmissionPage, error) {
	var resp SubmissionPage

	// 查询
	qc := params2Query(params)
	qc.Field.SelectId().SelectUserId().SelectProblemId().SelectStatus().SelectScore().SelectMemory().SelectTime().SelectLength().SelectLanguageId().SelectCreateTime().SelectUpdateTime()
	submissions, _, err := submission.Query.Select(qc)
	if err != nil {
		return resp, err
	}

	userIds := make([]int64, len(submissions))
	for i, s := range submissions {
		userIds[i] = s.UserId.Value()
	}
	problemIds := make([]int64, len(submissions))
	for i, s := range submissions {
		problemIds[i] = s.ProblemId.Value()
	}

	uqc := querycontext2.UserQueryContext{}
	uqc.Id.Add(userIds...)
	uqc.Field = *query.UserSimpleField
	users, _, err := user.Query.SelectByIds(uqc)

	pqc := querycontext2.ProblemQueryContext{}
	pqc.Id.Add(problemIds...)
	pqc.Field = *query.ProblemSimpleField
	problems, _, err := problem.Query.SelectByIds(pqc)

	for _, s := range submissions {
		respSubmission := domain2SubmissionData(s)

		// 获取用户信息
		if s.UserId.Value() != 0 {
			respSubmission.User = response.Domain2UserSimpleData(users[s.UserId.Value()])
		}

		// 获取题目信息
		if s.ProblemId.Value() != 0 {
			respSubmission.Problem = response.Domain2ProblemSimpleData(problems[s.ProblemId.Value()])
		}

		resp.Submissions = append(resp.Submissions, respSubmission)
	}

	resp.Page.Page = qc.Page.Page
	resp.Size = qc.Page.PageSize
	resp.Page.Total, err = Count(qc)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// SelectById 根据提交ID查询提交记录
func SelectById(sid int64, reqUser request.ReqUser) (response.RecordData, error) {
	var resp response.RecordData

	// 查询
	qc := querycontext2.SubmissionQueryContext{}
	qc.Id.Add(sid)
	qc.Field.SelectAll()
	s0, _, err := submission.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	// 隐藏源代码
	if reqUser.Role < entity.RoleAdmin && reqUser.Id != s0.UserId.Value() {
		s0.SourceCode.Set("")
	}

	// 获取评测结果
	jqc := querycontext2.JudgementQueryContext{}
	jqc.SubmissionId.Add(sid)
	jqc.Field.SelectAll()
	judgements, _, err := judgement.Query.Select(jqc)

	// 封装提交记录
	resp.Submission = domain2SubmissionData(s0)
	for _, j := range judgements {
		respJudgement := domain2JudgementData(j)

		resp.Judgements = append(resp.Judgements, respJudgement)
	}

	pqc := querycontext2.ProblemQueryContext{}
	pqc.Id.Add(s0.ProblemId.Value())
	pqc.Field = *query.ProblemSimpleField
	p, _, err := problem.Query.SelectOne(pqc)

	// 获取用户信息
	uqc := querycontext2.UserQueryContext{}
	uqc.Id.Add(s0.UserId.Value())
	uqc.Field = *query.UserSimpleField
	u, _, err := user.Query.SelectOne(uqc)

	resp.Submission.Problem = response.Domain2ProblemSimpleData(p)
	resp.Submission.User = response.Domain2UserSimpleData(u)
	return resp, nil
}

// SelectAcUsers 查询通过题目的用户
func SelectAcUsers(pid int64, size int64) ([]response.UserSimpleData, error) {
	var resp []response.UserSimpleData

	// 查询
	map_, err := submission.Query.SelectACUsers(pid, size)
	if err != nil {
		return resp, err
	}

	for _, m := range map_ {
		resp = append(resp, response.Map2UserSimpleData(m))
	}
	return resp, nil
}

func Statistics(params request.SubmissionStatisticsParams, reqUser request.ReqUser) (response.StatisticsRes, error) {
	qc := params2Query(params.QuerySubmissionParams)
	qc.GroupBy = params.GroupBy
	resp, err := submission.Query.GroupCount(qc)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
