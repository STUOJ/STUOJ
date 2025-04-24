package testcase

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/model"
)

type TestcasePage struct {
	Testcases []response.TestcaseData `json:"testcases"`
	model.Page
}

func Select(params request.QueryTestcaseParams, reqUser model.ReqUser) (TestcasePage, error) {
	var resp TestcasePage

	query_ := params2Query(params)
	query_.Field = *query.TestcaseAllField

	testcases, _, err := testcase.Query.Select(query_)
	if err != nil {
		return resp, err
	}

	resp.Testcases = make([]response.TestcaseData, len(testcases))
	for i, tc := range testcases {
		resp.Testcases[i] = domain2Resp(tc)
	}
	resp.Page.Page = query_.Page.Page
	resp.Page.Size = query_.Page.PageSize
	total, _ := GetStatistics(params)
	resp.Page.Total = total
	return resp, nil
}

// SelectById 根据ID查询评测点数据
func SelectById(id int64, reqUser model.ReqUser) (response.TestcaseData, error) {
	var resp response.TestcaseData

	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return resp, err
	}

	// 查询
	qc := querycontext.TestcaseQueryContext{}
	qc.Id.Add(id)
	qc.Field = *query.TestcaseAllField
	tc0, _, err := testcase.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(tc0)
	return resp, nil
}
