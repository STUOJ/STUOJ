package request

type QueryLanguageParams struct {
	Order   *string `form:"order,omitempty"`
	OrderBy *string `form:"order_by,omitempty"`
	Status  *int64  `form:"status,omitempty"`
}

type UpdateLanguageReq struct {
	ID     int64  `json:"id"`
	MapID  int64  `json:"map_id,omitempty"`
	Name   string `json:"name"`
	Serial int64  `json:"serial"`
	Status int64  `json:"status"`
}
