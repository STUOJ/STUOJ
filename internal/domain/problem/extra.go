package problem

import (
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
)

const (
	maxScoreSQL      = "(SELECT MAX(tbl_submission.score) FROM tbl_submission WHERE tbl_submission.problem_id = tbl_problem.id AND tbl_submission.user_id = %d) AS user_score"
	hasSubmissionSQL = "EXISTS (SELECT 1 FROM tbl_submission WHERE tbl_submission.problem_id = tbl_problem.id AND tbl_submission.user_id = %d) AS has_user_submission"
)

func QueryMaxScore(userId int64) option2.QueryContextOption {
	return func(pqm option2.QueryContext) option2.QueryContext {
		field := pqm.GetField()
		if field == nil {
			return pqm
		}
		maxScoreSelector := option2.NewSelector(maxScoreSQL, userId)
		hasUserSubmissionSelector := option2.NewSelector(hasSubmissionSQL, userId)
		field.AddSelect(*maxScoreSelector, *hasUserSubmissionSelector)
		return pqm
	}
}

const (
	QueryProblemTag = `(SELECT GROUP_CONCAT(DISTINCT tbl_problem_tag.tag_id) FROM tbl_problem_tag WHERE problem_id = tbl_problem.id) AS problem_tag_id`
)

func QueryTag() option2.QueryContextOption {
	return func(pqm option2.QueryContext) option2.QueryContext {
		field := pqm.GetField()
		selector := option2.NewSelector(QueryProblemTag)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	QueryProblemUser = `(SELECT GROUP_CONCAT(DISTINCT tbl_history.user_id) FROM tbl_history WHERE tbl_history.problem_id = tbl_problem.id) AS problem_user_id`
)

func QueryUser() option2.QueryContextOption {
	return func(pqm option2.QueryContext) option2.QueryContext {
		field := pqm.GetField()
		selector := option2.NewSelector(QueryProblemUser)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	WhereProblemTag = "tbl_problem.id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id In(?) GROUP BY problem_id HAVING COUNT(DISTINCT tag_id) =?)"
)

func WhereTag(tag []int64) option2.QueryContextOption {
	return func(pqm option2.QueryContext) option2.QueryContext {
		filter := pqm.GetExtraFilters()
		filter.Add(WhereProblemTag, option2.OpExtra, tag, len(tag))
		return pqm
	}
}

const (
	WhereProblemUser = "tbl_problem.id IN (SELECT tbl_history.problem_id FROM tbl_history WHERE tbl_history.user_id = ?)"
)

func WhereUser(userId int64) option2.QueryContextOption {
	return func(pqm option2.QueryContext) option2.QueryContext {
		filter := pqm.GetExtraFilters()
		filter.Add(WhereProblemUser, option2.OpExtra, userId)
		return pqm
	}
}
