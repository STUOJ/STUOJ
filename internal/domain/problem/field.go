package problem

import (
	"STUOJ/internal/db/query/option"
)

const (
	maxScoreSQL      = "(SELECT MAX(tbl_submission.score) FROM tbl_submission WHERE tbl_submission.problem_id = tbl_problem.id AND tbl_submission.user_id = %d) AS user_score"
	hasSubmissionSQL = "EXISTS (SELECT 1 FROM tbl_submission WHERE tbl_submission.problem_id = tbl_problem.id AND tbl_submission.user_id = %d) AS has_user_submission"
)

func SelectMaxScore(userId int64) option.QueryContextOption {
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
