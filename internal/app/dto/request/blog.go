package request

type QueryBlogParams struct {
	EndTime   *string `form:"end-time,omitempty"`
	Order     *string `form:"order,omitempty"`
	OrderBy   *string `form:"order_by,omitempty"`
	Page      *uint64 `form:"page,omitempty"`
	Problem   *string `form:"problem,omitempty"`
	Size      *uint64 `form:"size,omitempty"`
	StartTime *string `form:"start-time,omitempty"`
	Status    *string `form:"status,omitempty"`
	Title     *string `form:"title,omitempty"`
	User      *string `form:"user,omitempty"`
}

type CreateBlogReq struct {
	Content   string `json:"content"`
	ProblemId uint64 `json:"problem_id,omitempty"`
	Status    uint8  `json:"status"`
	Title     string `json:"title"`
	UserId    uint64 `json:"user_id"`
}

type UpdateBlogReq struct {
	Content   string `form:"content"`
	Id        uint64 `form:"id"`
	ProblemId uint64 `form:"problem_id,omitempty"`
	Status    uint8  `form:"status"`
	Title     string `form:"title"`
}
