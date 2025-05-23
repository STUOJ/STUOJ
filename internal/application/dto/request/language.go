package request

type QueryLanguageParams struct {
	Order   *string `form:"order,omitempty"`
	OrderBy *string `form:"order_by,omitempty"`
	Status  *string `form:"status,omitempty"`
}

type LanguageStatisticsParams struct {
	QueryLanguageParams
	GroupBy string `form:"group_by"`
}

type UpdateLanguageReq struct {
	Id     int64  `json:"id"`
	MapId  int64  `json:"map_id,omitempty"`
	Name   string `json:"name"`
	Serial int64  `json:"serial"`
	Status int64  `json:"status"`
}
