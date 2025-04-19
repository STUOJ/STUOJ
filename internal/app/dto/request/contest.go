package request

type QueryContestParams struct {
	EndTime   string `form:"end-time,omitempty"`
	Format    string `form:"format,omitempty"`
	Page      int64  `form:"page,omitempty"`
	Size      int64  `form:"size,omitempty"`
	StartTime string `form:"start-time,omitempty"`
	Status    string `form:"status,omitempty"`
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
