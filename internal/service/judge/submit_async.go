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
	"sync"
	"time"
)

func AsyncSubmit(s entity.Submission) (uint64, error) {
	var err error

	lang, err := language.SelectById(s.LanguageId)
	if err != nil {
		return 0, errors.New("获取语言信息失败")
	}
	if lang.Status != 3 {
		return 0, errors.New("该语言不可用")
	}

	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime
	s.Length = uint32(len(s.SourceCode))

	// 获取题目信息
	p, err := dao.SelectProblemById(s.ProblemId)
	if err != nil {
		log.Println(err)
		return 0, errors.New("获取题目信息失败")
	}

	// 获取评测点
	ts, err := dao.SelectTestcasesByProblemId(s.ProblemId)
	if err != nil || len(ts) < 1 {
		log.Println(err)
		return 0, errors.New("获取评测点数据失败")
	}

	// 插入提交
	s.Id, err = dao.InsertSubmission(s)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入提交信息失败")
	}

	// 异步提交
	asyncSubmit(s, p, ts)

	return s.Id, nil
}

// 异步提交
func asyncSubmit(s entity.Submission, p entity.Problem, ts []entity.Testcase) {
	if len(ts) == 0 {
		log.Println("No testcases provided")
		return
	}

	s.Status = entity.JudgeAC
	chJudgement := make(chan entity.Judgement, len(ts))
	var mu sync.Mutex
	var acCount uint64 = 0
	var wg sync.WaitGroup
	var errors []error

	lang, err := language.SelectById(s.LanguageId)
	if err != nil {
		log.Println("Failed to select language:", err)
		return
	}
	s1 := s
	s1.LanguageId = uint64(lang.MapId)

	wg.Add(len(ts))
	for _, t := range ts {
		go func(t entity.Testcase) {
			defer wg.Done()
			j, err := asyncJudge(s1, p, t)
			if err != nil {
				mu.Lock()
				errors = append(errors, err)
				j.Status = entity.JudgeIE
				mu.Unlock()
			}
			chJudgement <- j
		}(t)
	}
	wg.Wait()
	close(chJudgement)
	var judgements []entity.Judgement
	for j := range chJudgement {
		judgements = append(judgements, j)
		mu.Lock()
		s.Time = math.Max(s.Time, j.Time)
		s.Memory = max(s.Memory, j.Memory)
		if j.Status != entity.JudgeAC {
			if s.Status != entity.JudgeWA {
				s.Status = max(s.Status, j.Status)
			}
		} else {
			acCount++
		}
		mu.Unlock()
	}

	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		s.Status = entity.JudgeIE
	}

	if acCount > 0 {
		s.Score = uint8(100 * acCount / uint64(len(ts)))
	} else if (s.Status == entity.JudgeAC) || (s.Status == entity.JudgePD) || (s.Status == entity.JudgeIQ) {
		s.Status = entity.JudgeWA
	}

	s.UpdateTime = time.Now()

	err = dao.UpdateSubmissionByIdAndInsertJudgements(s, judgements)
	if err != nil {
		log.Println(err)
		return
	}
}

func asyncJudge(s entity.Submission, p entity.Problem, t entity.Testcase) (entity.Judgement, error) {
	var err error

	// 初始化评测点结果对象
	j := entity.Judgement{
		SubmissionId: s.Id,
		TestcaseId:   t.Id,
		Status:       entity.JudgePD,
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
