package request

type QueryBlogParams struct {
	EndTime   *string `form:"end-time,omitempty"`
	Order     *string `form:"order,omitempty"`
	OrderBy   *string `form:"order_by,omitempty"`
	Page      *int64  `form:"page,omitempty"`
	Problem   *string `form:"problem,omitempty"`
	Size      *int64  `form:"size,omitempty"`
	StartTime *string `form:"start-time,omitempty"`
	Status    *string `form:"status,omitempty"`
	Title     *string `form:"title,omitempty"`
	User      *string `form:"user,omitempty"`
}

type CreateBlogReq struct {
	Content   string `json:"content"`
	ProblemID int64  `json:"problem_id,omitempty"`
	Status    int64  `json:"status"`
	Title     string `json:"title"`
	UserID    int64  `json:"user_id"`
}

type UpdateBlogReq struct {
	Content   string `form:"content"`
	ID        int64  `form:"id"`
	ProblemID int64  `form:"problem_id,omitempty"`
	Status    int64  `form:"status"`
	Title     string `form:"title"`
}
