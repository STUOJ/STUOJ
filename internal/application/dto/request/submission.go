package request

type QuerySubmissionParams struct {
	EndTime   *string `form:"end-time,omitempty"`
	Language  *string `form:"language,omitempty"`
	Order     *string `form:"order,omitempty"`
	OrderBy   *string `form:"order_by,omitempty"`
	Page      *int64  `form:"page,omitempty"`
	Problem   *string `form:"problem,omitempty"`
	Size      *int64  `form:"size,omitempty"`
	StartTime *string `form:"start-time,omitempty"`
	Status    *string `form:"status,omitempty"`
	User      *string `form:"user,omitempty"`
}
