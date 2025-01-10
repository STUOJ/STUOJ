package judge

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/language"
	"errors"
	"log"
	"math"
	"strconv"
	"time"
)

// 等待提交
func WaitSubmit(s entity.Submission) (uint64, error) {
	var err error

	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime
	s.Length = uint64(len(s.SourceCode))

	// 获取题目信息
	p, err := dao.SelectProblemById(s.ProblemId)
	if err != nil {
		log.Println(err)
		return 0, errors.New("获取题目信息失败")
	}

	// 获取评测点
	ts, err := dao.SelectTestcasesByProblemId(s.ProblemId)
	if err != nil {
		log.Println(err)
		return 0, errors.New("获取评测点数据失败")
	}

	// 插入提交
	s.Id, err = dao.InsertSubmission(s)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入提交信息失败")
	}

	s.Status = entity.JudgeAC
	lang, err := language.SelectById(s.LanguageId)
	if err != nil {
		return 0, errors.New("获取语言信息失败")
	}
	if lang.Status != 3 {
		return 0, errors.New("该语言不可用")
	}
	s_ := s
	s_.LanguageId = lang.MapId

	// 提交评测点
	for _, t := range ts {
		j, err := waitJudge(s_, p, t)
		if err != nil {
			log.Println(err)
			s.Status = entity.JudgeIE
			continue
		}
		//log.Println(j)

		// 更新提交更新时间
		err = dao.UpdateSubmissionUpdateTimeById(j.SubmissionId)
		if err != nil {
			log.Println(err)
			s.Status = entity.JudgeIE
			continue
		}

		// 更新评测点结果
		err = dao.UpdateJudgementById(j)
		if err != nil {
			log.Println(err)
			s.Status = entity.JudgeIE
			continue
		}

		// 更新提交数据
		s.Time = math.Max(s.Time, j.Time)
		s.Memory = max(s.Memory, j.Memory)
		// 如果评测点结果不是AC，更新提交状态
		if j.Status != entity.JudgeAC {
			if s.Status != entity.JudgeWA {
				s.Status = max(s.Status, j.Status)
			}
		}
	}

	// 更新提交信息
	s.UpdateTime = time.Now()
	err = dao.UpdateSubmissionById(s)
	if err != nil {
		log.Println(err)
		return 0, errors.New("更新提交信息失败")
	}

	return s.Id, nil
}

// 等待评测
func waitJudge(s entity.Submission, p entity.Problem, t entity.Testcase) (entity.Judgement, error) {
	var err error

	// 初始化评测点结果对象
	j := entity.Judgement{
		SubmissionId: s.Id,
		TestcaseId:   t.Id,
		Status:       entity.JudgePD,
	}

	// 更新提交更新时间
	err = dao.UpdateSubmissionUpdateTimeById(j.SubmissionId)
	if err != nil {
		j.Status = entity.JudgeIE
		return j, err
	}

	// 插入评测点结果
	j.Id, err = dao.InsertJudgement(j)
	if err != nil {
		j.Status = entity.JudgeIE
		return j, err
	}

	// 初始化评测点评测对象
	judgeSubmission := model.JudgeSubmission{
		SourceCode:     s.SourceCode,
		LanguageId:     s.LanguageId,
		Stdin:          t.TestInput,
		ExpectedOutput: t.TestOutput,
		CPUTimeLimit:   p.TimeLimit,
		MemoryLimit:    p.MemoryLimit,
	}
	//log.Println(judgeSubmission)

	// 发送评测点评测请求（等待评测结果）
	result, err := judge0.Submit(judgeSubmission)
	if err != nil {
		log.Println(err)
		j.Status = entity.JudgeIE
		return j, err
	}
	//log.Println(result)

	// 解析时间
	time := float64(0)
	if result.Time != "" {
		time, err = strconv.ParseFloat(result.Time, 64)
		if err != nil {
			log.Println(err)
			j.Status = entity.JudgeIE
			return j, err
		}
	}

	// 更新评测点结果
	j.Time = time
	j.Memory = uint64(result.Memory)
	j.Stdout = result.Stdout
	j.Stderr = result.Stderr
	j.CompileOutput = result.CompileOutput
	j.Message = result.Message
	j.Status = entity.JudgeStatus(result.Status.Id)
	//log.Println(j)

	return j, nil
}
