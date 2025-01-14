package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 插入博客
func Insert(b entity.Blog) (uint64, error) {
	var err error

	updateTime := time.Now()
	b.UpdateTime = updateTime
	b.CreateTime = updateTime

	// 插入博客
	b.Id, err = dao.InsertBlog(b)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入博客失败")
	}

	return b.Id, nil
}

func BlogUpload(b entity.Blog, role entity.Role) (uint64, error) {
	var err error

	if b.Status == entity.BlogBanned || (b.Status == entity.BlogNotice && role < entity.RoleAdmin) {
		b.Status = entity.BlogDraft
	}

	updateTime := time.Now()
	b.UpdateTime = updateTime
	b.CreateTime = updateTime

	// 插入博客
	b.Id, err = dao.InsertBlog(b)
	if err != nil {
		log.Println(err)
		return 0, errors.New("保存博客失败")
	}

	return b.Id, nil
}
