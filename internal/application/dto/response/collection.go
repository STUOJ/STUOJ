package response

type CollectionData struct {
	Collaborator []UserSimpleData `json:"collaborator"`
	CreateTime   string           `json:"create_time"`
	Description  string           `json:"description"`
	Id           int64            `json:"id"`
	Problems     []struct {
		ProblemSimpleData
		ProblemUserScore
	} `json:"problems"`
	Status     int64          `json:"status"`
	Title      string         `json:"title"`
	UpdateTime string         `json:"update_time"`
	User       UserSimpleData `json:"user"`
}

type CollectionListItem struct {
	CreateTime string         `json:"create_time"`
	Id         int64          `json:"id"`
	Status     int64          `json:"status"`
	Title      string         `json:"title"`
	UpdateTime string         `json:"update_time"`
	User       UserSimpleData `json:"user"`
}
