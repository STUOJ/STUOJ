package collection

import "STUOJ/internal/db/query/option"

const (
	CollectionProblem = "tbl_collection.id IN (SELECT collection_id FROM tbl_collection_problem WHERE problem_id In(?) GROUP BY collection_id HAVING COUNT(DISTINCT problem_id) =?)"
)

func WhereProblem(ids []int64) option.QueryContextOption {
	return func(context option.QueryContext) option.QueryContext {
		filter := context.GetExtraFilters()
		filter.Add(CollectionProblem, option.OpExtra, ids, len(ids))
		return context
	}
}
