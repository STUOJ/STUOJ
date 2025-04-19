package testcase

import (
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

// SelectById 根据ID查询评测点数据
func SelectById(id uint64, reqUser model.ReqUser) (response.TestcaseData, error) {
	var resp response.TestcaseData

	// 检查权限
	if reqUser.Role < entity.RoleEditor {
		return resp, &errors.ErrUnauthorized
	}

	// 查询
	qc := querycontext.TestcaseQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId()
	tc0, _, err := testcase.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(tc0)
	return resp, nil
}
