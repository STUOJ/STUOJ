package handler

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"STUOJ/internal/service/blog"
	"STUOJ/internal/service/comment"
	"STUOJ/internal/service/judge"
	"STUOJ/internal/service/problem"
	"STUOJ/internal/service/record"
	"STUOJ/internal/service/tag"
	"STUOJ/internal/service/user"
	"STUOJ/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取用户统计信息
func StatisticsUser(c *gin.Context) {
	// 获取用户统计信息
	stats, err := user.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取用户角色统计信息
func StatisticsUserOfRole(c *gin.Context) {
	// 获取用户统计信息
	stats, err := user.GetStatisticsOfRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取用户注册统计信息
func StatisticsUserOfRegister(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取用户统计信息
	stats, err := user.GetStatisticsOfRegisterByPeriod(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取题目统计信息
func StatisticsProblem(c *gin.Context) {
	// 获取题目统计信息
	stats, err := problem.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取插入题目统计信息
func StatisticsProblemOfInsert(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取题目统计信息
	stats, err := problem.GetStatisticsOfInsertByPeriod(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取更新题目统计信息
func StatisticsProblemOfUpdate(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取题目统计信息
	stats, err := problem.GetStatisticsOfUpdateByPeriod(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取删除题目统计信息
func StatisticsProblemOfDelete(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取题目统计信息
	stats, err := problem.GetStatisticsOfDeleteByPeriod(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取标签统计信息
func StatisticsTag(c *gin.Context) {
	// 获取标签统计信息
	stats, err := tag.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取评测机统计信息
func StatisticsJudge(c *gin.Context) {
	statistics, err := judge.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", statistics))
}

// 获取提交记录提交信息
func StatisticsRecord(c *gin.Context) {

	// 获取提交记录统计信息
	stats, err := record.GetStatistics(dao.SubmissionWhere{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取提交语言统计信息
func StatisticsRecordOfLanguage(c *gin.Context) {
	// 获取提交记录统计信息
	stats, err := record.GetStatisticsOfSubmissionLanguage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取提交状态统计信息
func StatisticsSubmissionOfStatus(c *gin.Context) {
	// 获取提交记录统计信息
	stats, err := record.GetStatisticsOfSubmissionStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取评测状态统计信息
func StatisticsJudgementOfStatus(c *gin.Context) {
	// 获取提交记录统计信息
	stats, err := record.GetStatisticsOfJudgementStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取提交记录提交信息
func StatisticsRecordOfSubmit(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取提交记录统计信息
	stats, err := record.GetStatisticsOfSubmitByPeriod(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取博客统计信息
func StatisticsBlog(c *gin.Context) {
	// 获取博客统计信息
	stats, err := blog.GetStatistics(dao.BlogWhere{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取博客提交统计信息
func StatisticsBlogOfSubmit(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取博客统计信息
	stats, err := blog.GetStatisticsOfSubmitByPeriod(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取评论提交统计信息
func StatisticsCommentOfSubmit(c *gin.Context) {
	p, err := utils.GetPeriod(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.RespError(err.Error(), nil))
		return
	}

	// 获取评论统计信息
	stats, err := comment.GetStatisticsOfSubmitByPeriod(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}
