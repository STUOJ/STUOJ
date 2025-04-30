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
	ProblemId int64  `json:"problem_id,omitempty"`
	Status    uint8  `json:"status"`
	Title     string `json:"title"`
}

type UpdateBlogReq struct {
	Content   string `form:"content"`
	Id        int64  `form:"id"`
	ProblemId int64  `form:"problem_id,omitempty"`
	Status    uint8  `form:"status"`
	Title     string `form:"title"`
}
