package judge

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"log"
	"strconv"
)

func TestRun(s entity.Submission, stdin string) (entity.Judgement, error) {
	var err error
	j := entity.Judgement{}

	// 初始化评测点评测对象
	js := model.JudgeSubmission{
		SourceCode:     s.SourceCode,
		LanguageId:     s.LanguageId,
		Stdin:          stdin,
		ExpectedOutput: "",
		CPUTimeLimit:   2,
		MemoryLimit:    102400,
	}
	//log.Println(js)

	// 发送评测点评测请求（等待评测结果）
	result, err := judge0.Submit(js)
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
