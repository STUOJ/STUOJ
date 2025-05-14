package contest

import "STUOJ/internal/infrastructure/persistence/repository/option"

const (
	QueryContestProblem = "(SELECT GROUP_CONCAT(DISTINCT tbl_contest_problem.problem_id ORDER BY serial ASC) FROM tbl_contest_problem WHERE tbl_contest_problem.contest_id = tbl_contest.id) AS contest_problem_id"
)

func QueryProblemId() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		selector := option.NewSelector(QueryContestProblem)
		field.AddSelect(*selector)
		return pqm
	}
}

const (
	QueryContestUser = "(SELECT GROUP_CONCAT(DISTINCT tbl_contest_user.user_id ORDER BY tbl_contest_user.id ASC) FROM tbl_contest_user WHERE tbl_contest_user.contest_id = tbl_contest.id) AS contest_user_id"
)

func QueryUserId() option.QueryContextOption {
	return func(pqm option.QueryContext) option.QueryContext {
		field := pqm.GetField()
		selector := option.NewSelector(QueryContestUser)
		field.AddSelect(*selector)
		return pqm
	}
}
