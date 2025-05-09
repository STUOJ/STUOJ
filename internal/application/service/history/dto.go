package history

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/utils"
)

func params2Query(params request.QueryHistoryParams) querycontext.HistoryQueryContext {
	query := querycontext.HistoryQueryContext{}
	if params.ProblemId != nil {
		if problemIds, err := utils.StringToInt64Slice(*params.ProblemId); err == nil {
			query.ProblemId.Add(problemIds...)
		}
	}
	if params.UserId != nil {
		if userIds, err := utils.StringToInt64Slice(*params.UserId); err == nil {
			query.UserId.Add(userIds...)
		}
	}
	if params.Operation != nil {
		if operations, err := dao.StringToOperationSlice(*params.Operation); err == nil {
			query.Operation.Add(operations...)
		}
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option2.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option2.NewSortQuery(*params.OrderBy, *params.Order)
	}
	return query
}
