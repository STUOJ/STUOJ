package history

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/history"
	"STUOJ/internal/domain/user"
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
	userQuery.Id.Add(historyDomain.UserId)
	userQuery.Field = *query.UserSimpleField
	userDomain, _, err := user.Query.SelectOne(userQuery)
	if err == nil {
		res.User = response.Domain2UserSimpleData(userDomain)
	}
	return res, nil
}

/**
 * 查询历史记录列表
 * @param params 查询参数
 * @param reqUser 请求用户
 * @return 历史记录列表和分页信息
 */
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
		userQuery.Id.Add(v.UserId)
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
