package user

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/user"
)

// Count 统计用户数量
func Count(req request.QueryUserParams) (uint64, error) {
	query := params2Query(req)
	count, err := user.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return uint64(count), nil
}
