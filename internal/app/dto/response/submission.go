package response

type SubmissionData struct {
	CreateTime string            `json:"create_time"`
	ID         int64             `json:"id"`
	LanguageID int64             `json:"language_id"`
	Length     int64             `json:"length"`
	Memory     int64             `json:"memory,omitempty"`
	Problem    ProblemSimpleData `json:"problem"`
	Score      int64             `json:"score,omitempty"`
	SourceCode string            `json:"source_code,omitempty"`
	Status     int64             `json:"status"`
	Time       float64           `json:"time,omitempty"`
	UpdateTime string            `json:"update_time"`
	User       UserSimpleData    `json:"user"`
}
