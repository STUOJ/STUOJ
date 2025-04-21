package solution

import (
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/model"
)

// SelectById 根据ID查询评测点数据
func SelectById(id int64, reqUser model.ReqUser) (response.SolutionData, error) {
	var resp response.SolutionData

	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return resp, err
	}

	// 查询
	qc := querycontext.SolutionQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId()
	s0, _, err := solution.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(s0)
	return resp, nil
}
