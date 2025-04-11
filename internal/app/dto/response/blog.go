package response

type BlogData struct {
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	ID         int64  `json:"id"`
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
	ID        int64  `json:"id"`
	ProblemID *int64 `json:"problem_id,omitempty"`
	Title     string `json:"title"`
	UserID    int64  `json:"user_id"`
}
