package testcase

import "STUOJ/internal/domain/problem"

// UpdateProblemUpdateTime 更新题目更新时间
func updateProblemUpdateTime(id int64) error {
	p0 := problem.NewProblem(
		problem.WithId(id),
	)

	return p0.Update()
}
