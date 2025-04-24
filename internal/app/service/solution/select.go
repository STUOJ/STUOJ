package solution

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/model"
)

type SolutionPage struct {
	Solutions []response.SolutionData `json:"solutions"`
	model.Page
}

func Select(params request.QuerySolutionParams, reqUser model.ReqUser) (SolutionPage, error) {
	var res SolutionPage
	query_ := params2Query(params)
	query_.Field = *query.SolutionAllField
	solutions, _, err := solution.Query.Select(query_)
	if err != nil {
		return res, err
	}
	for _, s := range solutions {
		res.Solutions = append(res.Solutions, domain2Resp(s))
	}
	res.Page.Page = query_.Page.Page
	res.Page.Size = query_.Page.PageSize
	total, _ := GetStatistics(params)
	res.Page.Total = total
	return res, nil
}

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
	qc.Field = *query.SolutionAllField
	s0, _, err := solution.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(s0)
	return resp, nil
}
