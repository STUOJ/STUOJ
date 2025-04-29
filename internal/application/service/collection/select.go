package collection

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/user"
	entity2 "STUOJ/internal/infrastructure/repository/entity"
	query2 "STUOJ/internal/infrastructure/repository/query"
	querycontext2 "STUOJ/internal/infrastructure/repository/querycontext"
	model2 "STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
	"slices"
)

type CollectionPage struct {
	Collections []response.CollectionListItem `json:"collections"`
	model2.Page
}

// SelectById 根据Id查询题单
func SelectById(id int64, reqUser model2.ReqUser) (response.CollectionData, error) {
	var res response.CollectionData
	// 获取题单信息
	collectionQueryContext := querycontext2.CollectionQueryContext{}
	collectionQueryContext.Id.Add(id)
	collectionQueryContext.Field = *query2.CollectionAllField
	collectionDomain, collectionMap, err := collection.Query.SelectOne(collectionQueryContext, collection.QueryProblemId(), collection.QueryUserId())
	if err != nil {
		return res, err
	}
	if collectionDomain.Status < entity2.CollectionPublic {
		if err := isPermission(collectionDomain, reqUser); err != nil {
			return response.CollectionData{}, errors.ErrUnauthorized.WithMessage("没有权限查看该题单")
		}
	}
	res = domain2response(collectionDomain)

	problemQuery := querycontext2.ProblemQueryContext{}
	problemQuery.Field = *query2.ProblemSimpleField
	problemIds, _ := utils.StringToInt64Slice(collectionMap["collection_problem_id"].(string))
	problemQuery.Id.Set(problemIds)
	_, problemMap, err := problem.Query.SelectByIds(problemQuery, problem.QueryMaxScore(res.User.Id), problem.QueryTag())

	if err == nil {
		for _, p := range problemIds {
			problem_ := struct {
				response.ProblemSimpleData
				response.ProblemUserScore
			}{}
			problem_.ProblemSimpleData = response.Map2ProblemSimpleData(problemMap[int64(p)])
			res.Problems = append(res.Problems, problem_)
		}
	}

	userQuery := querycontext2.UserQueryContext{}
	userQuery.Field = *query2.UserSimpleField
	collaboratorIds, _ := utils.StringToInt64Slice(collectionMap["collection_user_id"].(string))
	userQuery.Id.Add(collectionDomain.UserId)
	userQuery.Id.Add(collaboratorIds...)
	userDomain, _, err := user.Query.SelectByIds(userQuery)

	if err == nil {
		for _, u := range collaboratorIds {
			res.Collaborator = append(res.Collaborator, response.Domain2UserSimpleData(userDomain[int64(u)]))
		}
		res.User = response.Domain2UserSimpleData(userDomain[int64(collectionDomain.UserId)])
	}
	return res, err
}

// Select 查询题单
func Select(params request.QueryCollectionParams, reqUser model2.ReqUser) (CollectionPage, error) {
	var res CollectionPage
	query_ := params2Model(params)
	if !query_.Status.Exist() {
		query_.Status.Set([]int64{int64(entity2.CollectionPublic)})
	}
	if slices.Contains(query_.Status.Value(), int64(entity2.CollectionPrivate)) && reqUser.Role < entity2.RoleAdmin {
		query_.UserId.Set([]int64{int64(reqUser.Id)})
	}
	query_.Field = *query2.CollectionListItemField
	collections, _, err := collection.Query.Select(query_)

	userIds := make([]int64, len(collections))
	for _, c := range collections {
		userIds = append(userIds, c.UserId)
	}

	userQuery := querycontext2.UserQueryContext{}
	userQuery.Field = *query2.UserSimpleField
	userQuery.Id.Set(userIds)

	users, _, err := user.Query.SelectByIds(userQuery)
	for _, collection_ := range collections {
		var resCollection response.CollectionListItem
		resCollection = domain2listItemResponse(collection_)
		resCollection.User = response.Domain2UserSimpleData(users[int64(collection_.UserId)])
		res.Collections = append(res.Collections, resCollection)
	}

	res.Page.Page = query_.Page.Page
	res.Page.Size = query_.Page.PageSize
	total, _ := GetStatistics(params)
	res.Page.Total = total
	return res, err
}
