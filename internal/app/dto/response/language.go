package response

type LanguageData struct {
	ID     int64  `json:"id"`
	MapID  int64  `json:"map_id,omitempty"`
	Name   string `json:"name"`
	Serial int64  `json:"serial"`
	Status int64  `json:"status"`
}
