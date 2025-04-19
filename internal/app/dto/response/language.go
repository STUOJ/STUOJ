package response

type LanguageData struct {
	Id     uint64 `json:"id"`
	MapId  uint32 `json:"map_id,omitempty"`
	Name   string `json:"name"`
	Serial uint16 `json:"serial"`
	Status uint8  `json:"status"`
}
