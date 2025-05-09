package response

type ContestData struct {
	Collaborator []UserSimpleData `json:"collaborator"`
	// 创建时间
	CreateTime  string `json:"create_time"`
	Description string `json:"description"`
	// 结束时间
	EndTime string `json:"end_time"`
	// 赛制
	Format int64 `json:"format"`
	// 比赛Id
	Id      int64 `json:"id"`
	Problem []struct {
		ProblemSimpleData
		ProblemUserScore
	} `json:"problem"`
	// 开始时间
	StartTime string `json:"start_time"`
	// 状态
	Status int64 `json:"status"`
	// 组队人数
	TeamSize int64  `json:"team_size"`
	Title    string `json:"title"`
	// 更新时间
	UpdateTime string         `json:"update_time"`
	User       UserSimpleData `json:"user"`
}

type ContestListItemData struct {
	Collaborator []UserSimpleData `json:"collaborator"`
	// 创建时间
	CreateTime string `json:"create_time"`
	// 结束时间
	EndTime string `json:"end_time"`
	// 赛制
	Format int64 `json:"format"`
	// 比赛Id
	Id int64 `json:"id"`
	// 开始时间
	StartTime string `json:"start_time"`
	// 状态
	Status int64 `json:"status"`
	// 组队人数
	TeamSize int64  `json:"team_size"`
	Title    string `json:"title"`
	// 更新时间
	UpdateTime string         `json:"update_time"`
	User       UserSimpleData `json:"user"`
}
