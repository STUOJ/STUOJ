package problem

import (
	"STUOJ/internal/infrastructure/persistence/entity"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

const (
	maxScoreSQL      = "(SELECT MAX(tbl_submission.score) FROM tbl_submission WHERE tbl_submission.problem_id = tbl_problem.id AND tbl_submission.user_id = %d) AS user_score"
	hasSubmissionSQL = "EXISTS (SELECT 1 FROM tbl_submission WHERE tbl_submission.problem_id = tbl_problem.id AND tbl_submission.user_id = %d) AS has_user_submission"
)

func QueryMaxScore(userId int64) option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		if field == nil {
			return pqm
		}
		maxScoreSelector := option.NewSelector(maxScoreSQL, userId)
		hasUserSubmissionSelector := option.NewSelector(hasSubmissionSQL, userId)
		field.AddSelect(*maxScoreSelector, *hasUserSubmissionSelector)
		return pqm
	}
}

const (
	QueryProblemTagSQL = `(SELECT GROUP_CONCAT(DISTINCT tbl_problem_tag.tag_id) FROM tbl_problem_tag WHERE problem_id = tbl_problem.id) AS problem_tag_id`
)

func QueryTag() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		selector := option.NewSelector(QueryProblemTagSQL)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	QueryProblemUserSQL = `(SELECT GROUP_CONCAT(DISTINCT tbl_history.user_id) FROM tbl_history WHERE tbl_history.problem_id = tbl_problem.id) AS problem_user_id`
)

func QueryUser() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		selector := option.NewSelector(QueryProblemUserSQL)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	WhereProblemTagSQL = "tbl_problem.id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id In(?) GROUP BY problem_id HAVING COUNT(DISTINCT tag_id) =?)"
)

func WhereTag(tag []int64) option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		filter := pqm.GetExtraFilters()
		filter.Add(WhereProblemTagSQL, option.OpExtra, tag, len(tag))
		return pqm
	}
}

const (
	WhereProblemUserSQL = "tbl_problem.id IN (SELECT tbl_history.problem_id FROM tbl_history WHERE tbl_history.user_id = ?)"
)

func WhereUser(userId int64) option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		filter := pqm.GetExtraFilters()
		filter.Add(WhereProblemUserSQL, option.OpExtra, userId)
		return pqm
	}
}

const (
	WhereUserNoACBeforeDateSQL = "tbl_problem.id NOT IN (SELECT DISTINCT  problem_id FROM `tbl_submission` WHERE user_id = ? AND status = ? AND DATE(create_time) < ?)"
)

func WhereUserNoACBeforeDate(userId int64, date time.Time) option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		dateStr := date.Format("2006-01-02")
		filter := pqm.GetExtraFilters()
		filter.Add(WhereUserNoACBeforeDateSQL, option.OpExtra, userId, entity.JudgeAC, dateStr)
		return pqm
	}
}
