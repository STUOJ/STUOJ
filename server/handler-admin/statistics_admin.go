package handler_admin

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
func AdminStatisticsUser(c *gin.Context) {
	// 获取用户统计信息
	stats, err := user.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取用户角色统计信息
func AdminStatisticsUserOfRole(c *gin.Context) {
	// 获取用户统计信息
	stats, err := user.GetStatisticsOfRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取用户注册统计信息
func AdminStatisticsUserOfRegister(c *gin.Context) {
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
func AdminStatisticsProblem(c *gin.Context) {
	// 获取题目统计信息
	stats, err := problem.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取插入题目统计信息
func AdminStatisticsProblemOfInsert(c *gin.Context) {
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
func AdminStatisticsProblemOfUpdate(c *gin.Context) {
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
func AdminStatisticsProblemOfDelete(c *gin.Context) {
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
func AdminStatisticsTag(c *gin.Context) {
	// 获取标签统计信息
	stats, err := tag.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取评测机统计信息
func AdminStatisticsJudge(c *gin.Context) {
	statistics, err := judge.GetStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", statistics))
}

// 获取提交记录提交信息
func AdminStatisticsRecord(c *gin.Context) {

	// 获取提交记录统计信息
	stats, err := record.GetStatistics(dao.SubmissionWhere{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取提交语言统计信息
func AdminStatisticsRecordOfLanguage(c *gin.Context) {
	// 获取提交记录统计信息
	stats, err := record.GetStatisticsOfSubmissionLanguage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取提交状态统计信息
func AdminStatisticsSubmissionOfStatus(c *gin.Context) {
	// 获取提交记录统计信息
	stats, err := record.GetStatisticsOfSubmissionStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取评测状态统计信息
func AdminStatisticsJudgementOfStatus(c *gin.Context) {
	// 获取提交记录统计信息
	stats, err := record.GetStatisticsOfJudgementStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取提交记录提交信息
func AdminStatisticsRecordOfSubmit(c *gin.Context) {
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
func AdminStatisticsBlog(c *gin.Context) {
	// 获取博客统计信息
	stats, err := blog.GetStatistics(dao.BlogWhere{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", stats))
}

// 获取博客提交统计信息
func AdminStatisticsBlogOfSubmit(c *gin.Context) {
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
func AdminStatisticsCommentOfSubmit(c *gin.Context) {
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
