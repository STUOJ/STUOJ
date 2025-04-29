package response

// RecordData 提交记录（提交信息+评测结果）
type RecordData struct {
	Submission SubmissionData  `json:"submission,omitempty"`
	Judgements []JudgementData `json:"judgements,omitempty"`
}
