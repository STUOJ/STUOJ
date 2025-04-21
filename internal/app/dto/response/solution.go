package response

type SolutionData struct {
	Id         int64  `json:"id"`
	LanguageId int64  `json:"language_id"`
	ProblemId  int64  `json:"problem_id"`
	SourceCode string `json:"source_code"`
}
