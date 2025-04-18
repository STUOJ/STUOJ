package response

type LanguageData struct {
	Id     int64  `json:"id"`
	MapId  int64  `json:"map_id,omitempty"`
	Name   string `json:"name"`
	Serial int64  `json:"serial"`
	Status int64  `json:"status"`
}
