package history

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/history"
	"STUOJ/internal/domain/user"
	query "STUOJ/internal/infrastructure/repository/query"
	querycontext "STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
)

type HistoryPage struct {
	Historys []response.HistoryListItemData `json:"historys"`
	model.Page
}

func SelectById(id int64, reqUser model.ReqUser) (response.HistoryData, error) {

	var res response.HistoryData
	// 创建查询选项
	historyQuery := querycontext.HistoryQueryContext{}
	historyQuery.Id.Add(id)
	historyQuery.Field = *query.HistoryAllField

	historyDomain, _, err := history.Query.SelectOne(historyQuery)
	if err != nil {
		return res, err
	}
	res = response.Domain2HistoryData(historyDomain)

	userQuery := querycontext.UserQueryContext{}
	userQuery.Id.Add(historyDomain.UserId.Value())
	userQuery.Field = *query.UserSimpleField
	userDomain, _, err := user.Query.SelectOne(userQuery)
	if err == nil {
		res.User = response.Domain2UserSimpleData(userDomain)
	}
	return res, nil
}

// Select 查询所有历史记录
func Select(params request.QueryHistoryParams, reqUser model.ReqUser) (HistoryPage, error) {
	var res HistoryPage
	historyQueryContext := params2Query(params)
	historyQueryContext.Field = *query.HistorySimpleField

	historyDomains, _, err := history.Query.Select(historyQueryContext)
	if err != nil {
		return HistoryPage{}, err
	}

	// 构建历史记录列表
	res.Historys = make([]response.HistoryListItemData, len(historyDomains))
	for i, v := range historyDomains {
		res.Historys[i] = response.Domain2HistoryListItem(v)

		// 获取用户信息
		userQuery := querycontext.UserQueryContext{}
		userQuery.Id.Add(v.UserId.Value())
		userQuery.Field = *query.UserSimpleField
		userDomain, _, err := user.Query.SelectOne(userQuery)
		if err == nil {
			res.Historys[i].User = response.Domain2UserSimpleData(userDomain)
		}
	}

	// 构建分页信息
	res.Page = model.Page{
		Page:  historyQueryContext.Page.Page,
		Size:  historyQueryContext.Page.PageSize,
		Total: int64(len(historyDomains)),
	}

	return res, nil
}
