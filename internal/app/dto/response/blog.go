package response

type BlogData struct {
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	Id         int64  `json:"id"`
	Problem    struct {
		ProblemSimpleData
		ProblemUserScore
	} `json:"problem,omitempty"`
	Status     int64          `json:"status"`
	Title      string         `json:"title"`
	UpdateTime string         `json:"update_time"`
	User       UserSimpleData `json:"user"`
}

type BlogSimpleData struct {
	Id        int64  `json:"id"`
	ProblemId *int64 `json:"problem_id,omitempty"`
	Title     string `json:"title"`
	UserId    int64  `json:"user_id"`
}
