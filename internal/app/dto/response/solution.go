package response

type SolutionData struct {
	ID         int64             `json:"id"`
	LanguageID int64             `json:"language_id"`
	Problem    ProblemSimpleData `json:"problem"`
	SourceCode string            `json:"source_code"`
}
