package request

type CreateTagReq struct {
	Name string `json:"name"`
}

type UpdateTagReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
