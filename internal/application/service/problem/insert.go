package problem

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/history"
	"STUOJ/internal/domain/problem"
	entity "STUOJ/internal/infrastructure/persistence/entity"
)

// Insert 创建新题目
// 创建题目信息，包括标题、描述、难度等基本信息以及关联的标签
func Insert(req request.CreateProblemReq, reqUser request.ReqUser) (int64, error) {
	// 创建新的题目对象
	p := problem.NewProblem(
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

	// 创建题目
	id, err := p.Create()
	if err != nil {
		return 0, err
	}

	// 更新题目标签
	p = problem.NewProblem(problem.WithId(id))
	err = p.UpdateTags(req.TagIds)
	if err != nil {
		return 0, err
	}

	// 记录操作历史
	h := history.NewHistory(
		history.WithUserId(reqUser.Id),
		history.WithProblemId(id),
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
		history.WithOperation(entity.OperationInsert),
	)
	_, err = h.Create()
	if err != nil {
		return 0, err
	}

	return id, nil
}
