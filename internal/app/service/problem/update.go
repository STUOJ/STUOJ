package problem

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/history"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/model"
)

// Update 根据ID更新题目
// 更新题目信息，包括标题、描述、难度等基本信息以及关联的标签
func Update(req request.UpdateProblemReq, reqUser model.ReqUser) error {
	// 查询题目
	qc := querycontext.ProblemQueryContext{}
	qc.Id.Add(req.Id)
	qc.Field.SelectAll()
	_, problemMap, err := problem.Query.SelectOne(qc, problem.QueryUser())
	if err != nil {
		return err
	}

	err = isPermission(problemMap, reqUser)
	if err != nil {
		return err
	}

	// 创建新的题目对象
	p1 := problem.NewProblem(
		problem.WithId(req.Id),
		problem.WithTitle(req.Title),
		problem.WithSource(req.Source),
		problem.WithDifficulty(entity.Difficulty(req.Difficulty)),
		problem.WithTimeLimit(float64(req.TimeLimit)),
		problem.WithMemoryLimit(req.MemoryLimit),
		problem.WithDescription(req.Description),
		problem.WithInput(req.Input),
		problem.WithOutput(req.Output),
		problem.WithSampleInput(req.SampleInput),
		problem.WithSampleOutput(req.SampleOutput),
		problem.WithHint(req.Hint),
		problem.WithStatus(entity.ProblemStatus(req.Status)),
	)

	// 更新题目基本信息
	err = p1.Update()
	if err != nil {
		return err
	}

	// 更新题目标签
	err = p1.UpdateTags(req.TagIds)
	if err != nil {
		return err
	}

	// 记录操作历史
	h := history.NewHistory(
		history.WithUserId(reqUser.Id),
		history.WithProblemId(req.Id),
		history.WithTitle(req.Title),
		history.WithSource(req.Source),
		history.WithDifficulty(entity.Difficulty(req.Difficulty)),
		history.WithTimeLimit(float64(req.TimeLimit)),
		history.WithMemoryLimit(req.MemoryLimit),
		history.WithDescription(req.Description),
		history.WithInput(req.Input),
		history.WithOutput(req.Output),
		history.WithSampleInput(req.SampleInput),
		history.WithSampleOutput(req.SampleOutput),
		history.WithHint(req.Hint),
		history.WithOperation(entity.OperationUpdate),
	)
	_, err = h.Create()
	if err != nil {
		return err
	}

	return nil
}
