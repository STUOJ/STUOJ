package problem

import (
	option "STUOJ/internal/infrastructure/persistence/repository/option"
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
	QueryProblemTag = `(SELECT GROUP_CONCAT(DISTINCT tbl_problem_tag.tag_id) FROM tbl_problem_tag WHERE problem_id = tbl_problem.id) AS problem_tag_id`
)

func QueryTag() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		selector := option.NewSelector(QueryProblemTag)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	QueryProblemUser = `(SELECT GROUP_CONCAT(DISTINCT tbl_history.user_id) FROM tbl_history WHERE tbl_history.problem_id = tbl_problem.id) AS problem_user_id`
)

func QueryUser() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		selector := option.NewSelector(QueryProblemUser)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	QueryContestMaxScoreSQL      = "(SELECT MAX(tbl_submission.score) FROM tbl_team_user INNER JOIN tbl_team ON tbl_team_user.team_id = tbl_team.id INNER JOIN tbl_team_submission ON tbl_team.id = tbl_team_submission.team_id INNER JOIN tbl_submission ON tbl_team_submission.submission_id = tbl_submission.id WHERE tbl_team.contest_id = %d AND tbl_team_user.user_id = %d AND tbl_submission.problem_id = tbl_problem.id) AS contest_user_score"
	QueryContestHasSubmissionSQL = "EXISTS (SELECT 1 FROM tbl_team_user INNER JOIN tbl_team ON tbl_team_user.team_id = tbl_team.id INNER JOIN tbl_team_submission ON tbl_team.id = tbl_team_submission.team_id INNER JOIN tbl_submission ON tbl_team_submission.submission_id = tbl_submission.id WHERE tbl_team.contest_id = %d AND tbl_team_user.user_id = %d AND tbl_submission.problem_id = tbl_problem.id) AS has_contest_submission"
)

func QueryContestMaxScore(contestId, userId int64) option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		maxScoreselector := option.NewSelector(QueryContestMaxScoreSQL, contestId, userId)
		hasSubmissionselector := option.NewSelector(QueryContestHasSubmissionSQL, contestId, userId)
		field.AddSelect(*maxScoreselector, *hasSubmissionselector)
		return pqm
	}
}

const (
	WhereProblemTag = "tbl_problem.id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id In(?) GROUP BY problem_id HAVING COUNT(DISTINCT tag_id) =?)"
)

func WhereTag(tag []int64) option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		filter := pqm.GetExtraFilters()
		filter.Add(WhereProblemTag, option.OpExtra, tag, len(tag))
		return pqm
	}
}

const (
	WhereProblemUser = "tbl_problem.id IN (SELECT tbl_history.problem_id FROM tbl_history WHERE tbl_history.user_id = ?)"
)

func WhereUser(userId int64) option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		filter := pqm.GetExtraFilters()
		filter.Add(WhereProblemUser, option.OpExtra, userId)
		return pqm
	}
}
