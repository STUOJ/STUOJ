package contest

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/contest"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/internal/infrastructure/persistence/repository/queryfield"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
	"slices"
	"time"
)

type ContestPage struct {
	Contests []response.ContestListItemData `json:"contests"`
	dto.Page
}

func Select(req request.QueryContestParams, reqUser request.ReqUser) (ContestPage, error) {
	var res ContestPage
	contestQuery := params2Query(req)
	contestQuery.Field = *queryfield.ContestSimpleField
	contestDomain, _, err := contest.Query.Select(contestQuery)
	if err != nil {
		return res, err
	}
	userIds := make([]int64, len(contestDomain))
	for i, c := range contestDomain {
		userId := c.UserId.Value()
		userIds[i] = userId
	}
	userQuery := querycontext.UserQueryContext{}
	userQuery.Field = *queryfield.UserSimpleField
	userQuery.Id.Set(userIds)
	users, _, err := user.Query.SelectByIds(userQuery)
	if err != nil {
		return res, err
	}
	for _, c := range contestDomain {
		var resContest response.ContestListItemData
		resContest = domain2listItemResponse(c)
		resContest.User = response.Domain2UserSimpleData(users[c.UserId.Value()])
		res.Contests = append(res.Contests, resContest)
	}
	res.Page.Page = contestQuery.Page.Page
	res.Page.Size = contestQuery.Page.PageSize
	total, _ := GetStatistics(req)
	res.Page.Total = total
	return res, nil
}

func SelectById(id int64, reqUser request.ReqUser) (response.ContestData, error) {
	var res response.ContestData
	contestQuery := querycontext.ContestQueryContext{}
	contestQuery.Field = *queryfield.ContestAllField
	contestQuery.Id.Set([]int64{id})
	contestDomain, contestMap, err := contest.Query.SelectOne(contestQuery, contest.QueryUserId(), contest.QueryProblemId(), contest.QueryJoinUserId())
	if err != nil {
		return res, err
	}
	var flag uint8 = 0 // 0 无权限 1 管理权限 2 参赛人员权限
	userIds, err := utils.StringToInt64Slice(string(contestMap["contest_user_id"].([]uint8)))
	if err != nil {
		return res, err
	}

	// 判断是否为比赛管理员
	if reqUser.Id == contestDomain.UserId.Value() || slices.Contains(userIds, reqUser.Id) || reqUser.Role >= entity.RoleAdmin {
		flag = 1
	}

	// 如果不是比赛管理员，判断是否为参赛人员
	if flag == 0 {
		joinUserIds, err := utils.StringToInt64Slice(string(contestMap["join_user_id"].([]uint8)))
		if err != nil {
			return res, err
		}
		if slices.Contains(joinUserIds, reqUser.Id) {
			flag = 2
		}
	}

	// 如果用户没有权限
	if flag == 0 {
		// 如果比赛不是公开的
		if contestDomain.Status.Value() < entity.ContestPublic {
			return res, errors.ErrUnauthorized.WithMessage("无权限查看")
		}
	}

	res = domain2response(contestDomain)

	// 统一处理用户信息查询（公开比赛也需要展示创建者信息）
	userQueryContext := querycontext.UserQueryContext{}
	userQueryContext.Id.Set(append(userIds, contestDomain.UserId.Value()))
	userQueryContext.Field = *queryfield.UserSimpleField
	userDomains, _, err := user.Query.SelectByIds(userQueryContext)
	if err != nil {
		return res, err
	}

	res.User = response.Domain2UserSimpleData(userDomains[contestDomain.UserId.Value()])
	res.Collaborator = make([]response.UserSimpleData, len(userIds))
	for i, v := range userIds {
		res.Collaborator[i] = response.Domain2UserSimpleData(userDomains[v])
	}

	// 如果用户没有权限或用户是参赛人员但比赛还未开始
	if flag == 0 || (flag == 2 && contestDomain.StartTime.Value().Before(time.Now())) {
		// 公开比赛直接返回
		return res, nil
	}
	// 处理题目信息查询
	problemIds, err := utils.StringToInt64Slice(string(contestMap["contest_problem_id"].([]uint8)))
	if err != nil {
		return res, err
	}
	problemQueryContext := querycontext.ProblemQueryContext{}
	problemQueryContext.Id.Set(problemIds)
	problemQueryContext.Field.SelectId().SelectTitle()
	_, problemMaps, err := problem.Query.SelectByIds(problemQueryContext, problem.QueryContestMaxScore(id, reqUser.Id))
	if err != nil {
		return res, err
	}
	for _, v := range problemIds {
		var problem_ struct {
			response.ProblemSimpleData
			response.ProblemUserScore
		}
		problem_.ProblemSimpleData = response.Map2ProblemSimpleData(problemMaps[v])
		problem_.ProblemUserScore = response.Map2ProblemUserScore(problemMaps[v])
		problem_.Id = v
		res.Problem = append(res.Problem, problem_)
	}
	return res, nil
}

func SelectProblem(contestId, problemSerial int64, reqUser request.ReqUser) (response.ProblemQueryData, error) {
	var res response.ProblemQueryData
	contsetQuery := querycontext.ContestQueryContext{}
	contsetQuery.Field.SelectId()
	_, contestMap, err := contest.Query.SelectOne(contsetQuery, contest.QueryProblemId())
	if err != nil {
		return res, err
	}
	problemIds, err := utils.StringToInt64Slice(string(contestMap["contest_problem_id"].([]uint8)))
	if err != nil {
		return res, err
	}
	problemId := problemIds[problemSerial-1]
	problemQuery := querycontext.ProblemQueryContext{}
	problemQuery.Field.SelectId().SelectTitle().SelectDescription().SelectInput().SelectOutput().SelectSampleInput().SelectSampleOutput().SelectTimeLimit().SelectMemoryLimit().SelectHint()
	problemQuery.Id.Add(problemId)
	problemDomain, problemMap, err := problem.Query.SelectOne(problemQuery, problem.QueryContestMaxScore(contestId, reqUser.Id))
	if err != nil {
		return res, err
	}
	res.ProblemData = response.Domain2ProblemData(problemDomain)
	res.ProblemUserScore = response.Map2ProblemUserScore(problemMap)
	return res, nil
}
