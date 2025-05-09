package request

type JudgeReq struct {
	// 语言id
	LanguageId int64 `json:"language_id"`
	// 题目id
	ProblemId int64 `json:"problem_id"`
	// 源程序代码
	SourceCode string `json:"source_code"`
}

type TestRunReq struct {
	LanguageId int64  `json:"language_id"`
	SourceCode string `json:"source_code"`
	Stdin      string `json:"stdin"`
}
