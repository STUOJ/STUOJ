package solution

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID更新题解
func Update(s entity.Solution) error {
	// 查询题解
	s0, err := dao.SelectSolutionById(s.Id)
	if err != nil {
		log.Println(err)
		return errors.New("获取题解失败")
	}

	s0.ProblemId = s.ProblemId
	s0.LanguageId = s.LanguageId
	s0.SourceCode = s.SourceCode

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(s.ProblemId)
	if err != nil {
		log.Println(err)
		return errors.New("更新题目更新时间失败")
	}

	// 更新题解
	err = dao.UpdateSolutionById(s0)
	if err != nil {
		return err
	}

	return nil
}
