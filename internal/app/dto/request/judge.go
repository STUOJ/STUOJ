package request

type JudgeReq struct {
	// 语言id
	LanguageID int64 `json:"language_id"`
	// 题目id
	ProblemID string `json:"problem_id"`
	// 源程序代码
	SourceCode string `json:"source_code"`
}

type TestRunReq struct {
	LanguageID int64  `json:"language_id"`
	SourceCode string `json:"source_code"`
	Stdin      string `json:"stdin"`
}
