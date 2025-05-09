package collection

import (
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

const (
	QueryCollectionProblem = "(SELECT GROUP_CONCAT(DISTINCT tbl_collection_problem.problem_id ORDER BY serial ASC) FROM tbl_collection_problem WHERE tbl_collection_problem.collection_id = tbl_collection.id) AS collection_problem_id"
)

func QueryProblemId() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		selector := option.NewSelector(QueryCollectionProblem)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	QueryCollectionUser = "(SELECT GROUP_CONCAT(DISTINCT tbl_collection_user.user_id) FROM tbl_collection_user WHERE tbl_collection_user.collection_id = tbl_collection.id) AS collection_user_id"
)

func QueryUserId() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		selector := option.NewSelector(QueryCollectionUser)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	WhereCollectionProblem = "tbl_collection.id IN (SELECT collection_id FROM tbl_collection_problem WHERE problem_id In(?) GROUP BY collection_id HAVING COUNT(DISTINCT problem_id) =?)"
)

func WhereProblem(ids []int64) option.QueryContextOption {
	return func(context option.QueryContext) option.QueryContext {
		filter := context.GetExtraFilters()
		filter.Add(WhereCollectionProblem, option.OpExtra, ids, len(ids))
		return context
	}
}
