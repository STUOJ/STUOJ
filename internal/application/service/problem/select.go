package problem

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/user"
	entity "STUOJ/internal/infrastructure/persistence/entity"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	querycontext2 "STUOJ/internal/infrastructure/persistence/repository/querycontext"
	query "STUOJ/internal/infrastructure/persistence/repository/queryfield"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
	"fmt"
	"hash/fnv"
	"slices"
	"time"
)

type ProblemPage struct {
	Problems []response.ProblemListItemData `json:"problems"`
	dto.Page
}

func SelectById(id int64, reqUser request.ReqUser) (response.ProblemQueryData, error) {
	var res response.ProblemQueryData
	problemQueryContext := querycontext2.ProblemQueryContext{}
	problemQueryContext.Id.Add(id)
	problemQueryContext.Field = *query.ProblemAllField
	problemDomain, problemMap, err := problem.Query.SelectOne(problemQueryContext, problem.QueryMaxScore(reqUser.Id), problem.QueryTag(), problem.QueryUser())
	if err != nil {
		return response.ProblemQueryData{}, err
	}
	res.ProblemData = domain2response(problemDomain)
	res.ProblemUserScore = response.Map2ProblemUserScore(problemMap)
	res.TagIds = response.Map2TagIds(problemMap)

	userIds, err := utils.StringToInt64Slice(string(problemMap["problem_user_id"].([]uint8)))
	if err != nil {
		return response.ProblemQueryData{}, errors.ErrInternalServer.WithMessage("获取题目修改者id失败")
	}

	if problemDomain.Status.Value() < entity.ProblemPublic && reqUser.Role < entity.RoleAdmin && !slices.Contains(userIds, reqUser.Id) {
		return response.ProblemQueryData{}, errors.ErrUnauthorized.WithMessage("无权限查看")
	}

	userQueryContext := querycontext2.UserQueryContext{}
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

func Select(params request.QueryProblemParams, reqUser request.ReqUser) (ProblemPage, error) {
	var res ProblemPage
	problemQueryContext := params2Query(params)

	if !problemQueryContext.Status.Exist() {
		problemQueryContext.Status.Add(entity.ProblemPublic)
	} else if slices.ContainsFunc(problemQueryContext.Status.Value(), func(s entity.ProblemStatus) bool { return s != entity.ProblemPublic }) && reqUser.Role < entity.RoleAdmin {
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
	res.Page = dto.Page{
		Page: problemQueryContext.Page.Page,
		Size: problemQueryContext.Page.PageSize,
	}
	total, _ := Count(problemQueryContext)
	res.Page.Total = total
	return res, nil
}

func Statistics(params request.ProblemStatisticsParams, reqUser request.ReqUser) (response.StatisticsRes, error) {
	problemQueryContext := params2Query(params.QueryProblemParams)
	problemQueryContext.GroupBy = params.GroupBy
	resp, err := problem.Query.GroupCount(problemQueryContext)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func SelectDailyProblem(reqUser request.ReqUser) (response.ProblemSimpleWithUserScore, error) {
	date := time.Now()
	dateStr := date.Format("2006-01-02")

	problemQueryContext := querycontext2.ProblemQueryContext{}

	problemQueryContext.Status.Add(entity.ProblemPublic)

	hasher := fnv.New64()
	hasher.Write([]byte(dateStr))
	hasher.Write([]byte{0}) // 添加分隔符
	hasher.Write([]byte(fmt.Sprintf("%d", reqUser.Id)))
	hashValue := hasher.Sum64()

	total, err := problem.Query.Count(problemQueryContext, problem.WhereUserNoACBeforeDate(reqUser.Id, date))
	if err != nil {
		return response.ProblemSimpleWithUserScore{}, err
	}

	index := hashValue % uint64(total)

	problemQueryContext.Field = *query.ProblemSimpleField
	problemQueryContext.Page = option.NewPagination(int64(index), 1)
	problemDomain, problemMap, err := problem.Query.Select(problemQueryContext, problem.QueryMaxScore(reqUser.Id), problem.WhereUserNoACBeforeDate(reqUser.Id, date))
	if err != nil {
		return response.ProblemSimpleWithUserScore{}, err
	}
	var res response.ProblemSimpleWithUserScore
	res.ProblemSimpleData = response.Domain2ProblemSimpleData(problemDomain[0])
	res.ProblemUserScore = response.Map2ProblemUserScore(problemMap[0])
	return res, nil
}
