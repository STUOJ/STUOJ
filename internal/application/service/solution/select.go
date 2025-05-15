package solution

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/internal/infrastructure/persistence/repository/queryfield"
)

type SolutionPage struct {
	Solutions []response.SolutionData `json:"solutions"`
	dto.Page
}

func Select(params request.QuerySolutionParams, reqUser request.ReqUser) (SolutionPage, error) {
	var res SolutionPage
	query_ := params2Query(params)
	query_.Field = *queryfield.SolutionAllField
	solutions, _, err := solution.Query.Select(query_)
	if err != nil {
		return res, err
	}
	for _, s := range solutions {
		res.Solutions = append(res.Solutions, domain2Resp(s))
	}
	res.Page.Page = query_.Page.Page
	res.Page.Size = query_.Page.PageSize
	total, _ := Count(query_)
	res.Page.Total = total
	return res, nil
}

// SelectById 根据ID查询评测点数据
func SelectById(id int64, reqUser request.ReqUser) (response.SolutionData, error) {
	var resp response.SolutionData

	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return resp, err
	}

	// 查询
	qc := querycontext.SolutionQueryContext{}
	qc.Id.Add(id)
	qc.Field = *queryfield.SolutionAllField
	s0, _, err := solution.Query.SelectOne(qc)
	if err != nil {
		return resp, err
	}

	resp = domain2Resp(s0)
	return resp, nil
}

func Statistics(params request.SolutionStatisticsParams, reqUser request.ReqUser) (response.StatisticsRes, error) {
	qc := params2Query(params.QuerySolutionParams)
	qc.GroupBy = params.GroupBy
	resp, err := solution.Query.GroupCount(qc)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
