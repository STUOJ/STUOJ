package request

type QueryTagParams struct {
	Name *string `json:"name,omitempty"`
	Page *int64  `json:"page,omitempty"`
	Size *int64  `json:"size,omitempty"`
}

type CreateTagReq struct {
	Name string `json:"name"`
}

type UpdateTagReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
