package problem

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/user"
	entity "STUOJ/internal/infrastructure/repository/entity"
	query "STUOJ/internal/infrastructure/repository/query"
	querycontext "STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
	"slices"
)

type ProblemPage struct {
	Problems []response.ProblemListItemData `json:"problems"`
	model.Page
}

func SelectById(id int64, reqUser model.ReqUser) (response.ProblemQueryData, error) {
	var res response.ProblemQueryData
	problemQueryContext := querycontext.ProblemQueryContext{}
	problemQueryContext.Id.Add(id)
	problemQueryContext.Field = *query.ProblemAllField
	problemDomain, problemMap, err := problem.Query.SelectOne(problemQueryContext, problem.QueryMaxScore(reqUser.Id), problem.QueryTag(), problem.QueryUser())
	if err != nil {
		return response.ProblemQueryData{}, err
	}
	res.ProblemData = domain2response(problemDomain)
	res.ProblemUserScore = response.Map2ProblemUserScore(problemMap)
	res.TagIds = response.Map2TagIds(problemMap)

	userIds, err := utils.StringToInt64Slice(problemMap["problem_user_id"].(string))
	if err != nil {
		return response.ProblemQueryData{}, errors.ErrInternalServer.WithMessage("获取题目修改者id失败")
	}

	if problemDomain.Status.Value() < entity.ProblemPublic && reqUser.Role < entity.RoleAdmin && !slices.Contains(userIds, reqUser.Id) {
		return response.ProblemQueryData{}, errors.ErrUnauthorized.WithMessage("无权限查看")
	}

	userQueryContext := querycontext.UserQueryContext{}
	userQueryContext.Id.Set(userIds)
	userQueryContext.Field = *query.UserSimpleField
	userDomain, _, err := user.Query.Select(userQueryContext)
	if err != nil {
		return response.ProblemQueryData{}, err
	}
	for _, v := range userDomain {
		res.User = append(res.User, response.Domain2UserSimpleData(v))
	}
	return res, nil
}

func Select(params request.QueryProblemParams, reqUser model.ReqUser) (ProblemPage, error) {
	var res ProblemPage
	problemQueryContext := params2Query(params)

	if !problemQueryContext.Status.Exist() {
		problemQueryContext.Status.Add(entity.ProblemPublic)
	} else if len(slices.DeleteFunc(problemQueryContext.Status.Value(), func(s entity.ProblemStatus) bool { return s == entity.ProblemPublic })) > 0 && reqUser.Role < entity.RoleAdmin {
		problem.WhereUser(reqUser.Id)(&problemQueryContext)
	}

	problemQueryContext.Field = *query.ProblemListItemField

	problemDomain, problemMap, err := problem.Query.Select(problemQueryContext, problem.QueryMaxScore(reqUser.Id), problem.QueryTag(), problem.QueryUser())
	if err != nil {
		return ProblemPage{}, err
	}
	res.Problems = make([]response.ProblemListItemData, len(problemDomain))
	for i, v := range problemDomain {
		res.Problems[i] = response.Domain2ProblemListItemData(v)
		res.Problems[i].ProblemUserScore = response.Map2ProblemUserScore(problemMap[i])
		res.Problems[i].TagIds = response.Map2TagIds(problemMap[i])
	}
	res.Page = model.Page{
		Page: problemQueryContext.Page.Page,
		Size: problemQueryContext.Page.PageSize,
	}
	total, _ := GetStatistics(params)
	res.Page.Total = total
	return res, nil
}

func Statistics(params request.ProblemStatisticsParams, reqUser model.ReqUser) (response.StatisticsRes, error) {
	problemQueryContext := params2Query(params.QueryProblemParams)
	problemQueryContext.GroupBy = params.GroupBy
	resp, err := problem.Query.GroupCount(problemQueryContext)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
