package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

func Update(b entity.Blog, userId uint64, role entity.Role) error {
	// 查询博客
	b0, err := dao.SelectBlogById(b.Id)
	if err != nil {
		log.Println(err)
		return errors.New("博客不存在")
	}

	if role < entity.RoleAdmin {
		if b0.UserId != userId {
			return errors.New("没有权限，只能编辑自己的博客")
		}
		if b0.Status == entity.BlogBanned {
			return errors.New("该博客已被封禁，无法编辑")
		}
		if b.Status == entity.BlogBanned {
			return errors.New("只有管理员才能封禁博客")
		}
		if b.Status == entity.BlogNotice {
			return errors.New("只有管理员才能发布公告")
		}
	}

	updateTime := time.Now()
	b0.Title = b.Title
	b0.Content = b.Content
	b0.UpdateTime = updateTime
	b0.Status = b.Status
	b0.ProblemId = b.ProblemId

	// 更新博客
	err = dao.UpdateBlogById(b0)
	if err != nil {
		log.Println(err)
		return errors.New("修改博客失败")
	}

	return nil
}
