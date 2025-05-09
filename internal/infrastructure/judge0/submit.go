package judge0

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func Submit(submission JudgeSubmission) (JudgeResult, error) {
	data, err := json.Marshal(submission)
	if err != nil {
		return JudgeResult{}, err
	}
	bodystr, err := httpInteraction("/submissions", "POST", bytes.NewReader(data))
	if err != nil {
		return JudgeResult{}, err
	}
	var result JudgeResult
	err = json.Unmarshal([]byte(bodystr), &result)
	if err != nil {
		return JudgeResult{}, err
	}

	return result, nil
}

func QueryResult(token string) (JudgeResult, error) {
	bodystr, err := httpInteraction("/submissions"+"/"+token, "GET", nil)
	if err != nil {
		return JudgeResult{}, err
	}
	var result JudgeResult
	err = json.Unmarshal([]byte(bodystr), &result)
	if err != nil {
		return JudgeResult{}, err
	}
	return result, nil
}

func QueryResults(page uint64, per_page uint64) (JudgeSubmissionResults, error) {
	bodystr, err := httpInteraction("/submissions"+"/?page="+strconv.FormatUint(page, 10)+"&per_page="+strconv.FormatUint(per_page, 10), "GET", nil)
	if err != nil {
		return JudgeSubmissionResults{}, err
	}
	var results JudgeSubmissionResults
	err = json.Unmarshal([]byte(bodystr), &results)
	if err != nil {
		return JudgeSubmissionResults{}, err
	}
	return results, nil
}
