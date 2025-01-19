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
func SelectById(id uint64, userId uint64, role entity.Role) (entity.Blog, error) {
	b, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return entity.Blog{}, errors.New("获取博客失败")
	}
	if b.Status != entity.BlogPublic && role < entity.RoleAdmin && b.UserId != userId {
		return entity.Blog{}, errors.New("该博客未公开")
	}
	return b, nil
}

func Select(condition model.BlogWhere, userId uint64, role entity.Role) (BlogPage, error) {
	if !condition.Status.Exist() {
		condition.Status.Set(entity.BlogPublic)
	} else {
		for _, v := range condition.Status.Value() {
			if entity.BlogStatus(v) < entity.BlogPublic {
				if role < entity.RoleAdmin {
					condition.Status.Set(entity.BlogPublic)
				}
			}
		}
	}
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	blogs, err := dao.SelectBlogs(condition)
	if err != nil {
		log.Println(err)
		return BlogPage{}, errors.New("获取博客失败")
	}
	count, err := dao.CountBlogs(condition)

	cutBlogContent(blogs)

	if err != nil {
		log.Println(err)
		return BlogPage{}, errors.New("获取统计失败")
	}
	bPage := BlogPage{
		Blogs: blogs,
		Page: model.Page{
			Total: count,
			Size:  condition.Size.Value(),
			Page:  condition.Page.Value(),
		},
	}
	return bPage, nil
}

// 不返回正文
func cutBlogContent(blogs []entity.Blog) {
	for i := range blogs {
		if len(blogs[i].Content) > 256 {
			blogs[i].Content = blogs[i].Content[:256] + "..."
		}
	}
}
