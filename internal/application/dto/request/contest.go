package request

type QueryContestParams struct {
	BeginEnd    *string `json:"begin-end,omitempty"`
	BeginStart  *string `json:"begin-start,omitempty"`
	EndTime     *string `json:"end-time,omitempty"`
	FinishEnd   *string `json:"finish-end,omitempty"`
	FinishStart *string `json:"finish-start,omitempty"`
	Format      *string `json:"format,omitempty"`
	TeamSize    *string `json:"team-size,omitempty"`
	Order       *string `json:"order,omitempty"`
	OrderBy     *string `json:"order-by,omitempty"`
	Page        *int64  `json:"page,omitempty"`
	Size        *int64  `json:"size,omitempty"`
	StartTime   *string `json:"start-time,omitempty"`
	Status      *string `json:"status,omitempty"`
	Title       *string `json:"title,omitempty"`
	UserId      *int64  `json:"user_id,omitempty"`
}

type ContestStatisticsParams struct {
	QueryContestParams
	GroupBy string `form:"group-by,omitempty"`
}

type CreateContestReq struct {
	Description string `form:"description"`
	// 结束时间
	EndTime string `form:"end_time"`
	// 赛制
	Format int64 `form:"format"`
	// 开始时间
	StartTime string `form:"start_time"`
	// 状态
	Status int64 `form:"status"`
	// 组队人数
	TeamSize   int64  `form:"team_size"`
	Title      string `form:"title"`
	Collection int64  `form:"collection,omitempty"`
}

type UpdateContestReq struct {
	Description string `form:"description"`
	// 结束时间
	EndTime string `form:"end_time"`
	// 赛制
	Format int64 `form:"format"`
	// 比赛Id
	Id int64 `form:"id"`
	// 开始时间
	StartTime string `form:"start_time"`
	// 状态
	Status int64 `form:"status"`
	// 组队人数
	TeamSize int64  `form:"team_size"`
	Title    string `form:"title"`
}
