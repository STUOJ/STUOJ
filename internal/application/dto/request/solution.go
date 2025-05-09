package request

type QuerySolutionParams struct {
	Language *string `json:"language,omitempty"`
	Page     *int64  `json:"page,omitempty"`
	Problem  *string `json:"problem,omitempty"`
	Size     *int64  `json:"size,omitempty"`
}

type SolutionStatisticsParams struct {
	QuerySolutionParams
	GroupBy string `form:"group_by"`
}

type CreateSolutionReq struct {
	LanguageId int64  `json:"language_id"`
	ProblemId  int64  `json:"problem_id"`
	SourceCode string `json:"source_code"`
}

type UpdateSolutionReq struct {
	Id         int64  `json:"id"`
	LanguageId int64  `json:"language_id"`
	ProblemId  int64  `json:"problem_id"`
	SourceCode string `json:"source_code"`
}
