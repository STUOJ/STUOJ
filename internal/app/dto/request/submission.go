package request

type QuerySubmissionParams struct {
	Distinct       *string `form:"distinct,omitempty"`
	EndTime        *string `form:"end-time,omitempty"`
	ExcludeHistory *bool   `form:"exclude_history,omitempty"`
	Language       *int64  `form:"language,omitempty"`
	Order          *string `form:"order,omitempty"`
	OrderBy        *string `form:"order_by,omitempty"`
	Page           *int64  `form:"page,omitempty"`
	Problem        *int64  `form:"problem,omitempty"`
	Size           *int64  `form:"size,omitempty"`
	StartTime      *string `form:"start-time,omitempty"`
	Status         *int64  `form:"status,omitempty"`
	User           *int64  `form:"user,omitempty"`
}
