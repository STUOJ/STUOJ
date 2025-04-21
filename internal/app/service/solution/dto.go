package solution

import (
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/domain/solution"
)

func domain2Resp(dm solution.Solution) response.SolutionData {
	return response.SolutionData{
		Id:         dm.Id,
		ProblemId:  dm.ProblemId,
		LanguageId: dm.LanguageId,
		SourceCode: dm.SourceCode.String(),
	}
}
