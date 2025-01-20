package comment

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 根据ID更新评论
func Update(c entity.Comment) error {
	// 查询评论
	c0, err := dao.SelectCommentById(c.Id)
	if err != nil {
		log.Println(err)
		return errors.New("评论不存在")
	}

	updateTime := time.Now()
	c0.UpdateTime = updateTime
	c0.UserId = c.UserId
	c0.BlogId = c.BlogId
	c0.Content = c.Content
	c0.Status = c.Status

	// 更新评论
	err = dao.UpdateCommentById(c0)
	if err != nil {
		log.Println(err)
		return errors.New("更新评论失败")
	}

	return nil
}
