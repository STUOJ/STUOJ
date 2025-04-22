package request

type QueryTagParams struct {
	Id      *string `form:"id,omitempty"`
	Name    *string `form:"name,omitempty"`
	Order   *string `form:"order,omitempty"`
	OrderBy *string `form:"order_by,omitempty"`
	Page    *int64  `form:"page,omitempty"`
	Size    *int64  `form:"size,omitempty"`
}

type CreateTagReq struct {
	Name string `json:"name"`
}

type UpdateTagReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
