package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

type BlogPage struct {
	Blogs []entity.Blog `json:"blogs"`
	model.Page
}

// 根据ID查询博客
func SelectById(id uint64, userId uint64, admin ...bool) (entity.Blog, error) {
	b, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return entity.Blog{}, errors.New("获取博客失败")
	}
	if b.Status != entity.BlogStatusPublic && (len(admin) == 0 || !admin[0]) && b.UserId != userId {
		return entity.Blog{}, errors.New("该博客未公开")
	}
	return b, nil
}

func Select(condition dao.BlogWhere, userId uint64, page uint64, size uint64, admin ...bool) (BlogPage, error) {
	if len(admin) == 0 || !admin[0] && (!condition.UserId.Exist() || condition.UserId.Value() != userId) {
		condition.Status.Set(entity.BlogStatusPublic)
	}
	blogs, err := dao.SelectBlogs(condition, page, size)
	if err != nil {
		log.Println(err)
		return BlogPage{}, errors.New("获取博客失败")
	}
	count, err := dao.CountBlogs(condition)

	hideBlogContent(blogs)

	if err != nil {
		log.Println(err)
		return BlogPage{}, errors.New("获取统计失败")
	}
	bPage := BlogPage{
		Blogs: blogs,
		Page: model.Page{
			Total: count,
			Size:  size,
			Page:  page,
		},
	}
	return bPage, nil
}

// 不返回正文
func hideBlogContent(blogs []entity.Blog) {
	for i := range blogs {
		blogs[i].Content = ""
	}
}
