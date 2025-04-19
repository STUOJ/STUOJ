package response

type SolutionData struct {
	Id         int64             `json:"id"`
	LanguageId int64             `json:"language_id"`
	Problem    ProblemSimpleData `json:"problem"`
	SourceCode string            `json:"source_code"`
}
