package contest

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/contest"
)

func GetStatistics(params request.QueryContestParams) (int64, error) {
	query := params2Query(params)
	return contest.Query.Count(query)
}
