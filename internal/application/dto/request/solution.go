package request

type QuerySolutionParams struct {
	Language *string `form:"language,omitempty"`
	Page     *int64  `form:"page,omitempty"`
	Problem  *string `form:"problem,omitempty"`
	Size     *int64  `form:"size,omitempty"`
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
