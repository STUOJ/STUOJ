package tag

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/tag"
	"STUOJ/utils"
)

func domain2Resp(dm tag.Tag) response.TagData {
	return response.TagData{
		Id:   dm.Id,
		Name: dm.Name.String(),
	}
}

func params2Query(params request.QueryTagParams) (query querycontext.TagQueryContext) {
	if params.Id != nil {
		ids, err := utils.StringToInt64Slice(*params.Id)
		if err != nil {
			query.Id.Set(ids)
		}
	}
	if params.Name != nil {
		query.Name.Add(*params.Name)
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.Order, *params.OrderBy)
	}
	return query
}
