package solution

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/solution"
)

// 插入题解
func Insert(req request.CreateSolutionReq, reqUser request.ReqUser) (int64, error) {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return 0, err
	}

	s1 := solution.NewSolution(
		solution.WithProblemId(req.ProblemId),
		solution.WithLanguageId(req.LanguageId),
		solution.WithSourceCode(req.SourceCode),
	)

	// 更新题目更新时间
	err = updateProblemUpdateTime(req.ProblemId)
	if err != nil {
		return 0, err
	}

	return s1.Create()
}
