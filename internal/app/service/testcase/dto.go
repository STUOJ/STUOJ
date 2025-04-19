package testcase

import (
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/domain/testcase"
)

func domain2Resp(dm testcase.Testcase) response.TestcaseData {
	return response.TestcaseData{
		Id:         dm.Id,
		ProblemId:  dm.ProblemId,
		Serial:     dm.Serial,
		TestInput:  dm.TestInput.String(),
		TestOutput: dm.TestOutput.String(),
	}
}
