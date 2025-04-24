package request

type QueryHistoryParams struct {
	Operation *string `json:"operation,omitempty"`
	Order     *string `json:"order,omitempty"`
	OrderBy   *string `json:"order_by,omitempty"`
	Page      *int64  `json:"page,omitempty"`
	ProblemId *string `json:"problem_id,omitempty"`
	Size      *int64  `json:"size,omitempty"`
	UserId    *string `json:"user_id,omitempty"`
}
