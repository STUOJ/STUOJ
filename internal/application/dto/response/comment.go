package response

type CommentData struct {
	Blog       BlogSimpleData `json:"blog"`
	Content    string         `json:"content"`
	CreateTime string         `json:"create_time"`
	Id         int64          `json:"id"`
	Status     int64          `json:"status"`
	UpdateTime string         `json:"update_time"`
	User       UserSimpleData `json:"user"`
}
