package request

type QueryCommentParams struct {
	User      *int64  `form:"user,omitempty"`
	Blog      *int64  `form:"blog,omitempty"`
	EndTime   *string `form:"end-time,omitempty"`
	Order     *string `form:"order,omitempty"`
	OrderBy   *string `form:"order_by"`
	Page      *int64  `form:"page,omitempty"`
	Size      *int64  `form:"size,omitempty"`
	StartTime *string `form:"start-time,omitempty"`
	Status    *string `form:"status,omitempty"`
}

type CommentStatisticsParams struct {
	QueryCommentParams
	GroupBy string `form:"group_by"`
}

type CreateCommentReq struct {
	BlogId  int64  `json:"blog_id"`
	Content string `json:"content"`
}

type UpdateCommentReq struct {
	Content string `json:"content"`
	Id      int64  `json:"id"`
	Status  int64  `json:"status"`
}
