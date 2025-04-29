package blog

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/blog"
	entity2 "STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

func Update(req request.UpdateBlogReq, reqUser model.ReqUser) error {
	blogQueryContext := querycontext.BlogQueryContext{}
	blogQueryContext.Id.Add(req.Id)
	blogQueryContext.Field.SelectId().SelectStatus().SelectUserId()
	b0, _, err := blog.Query.SelectOne(blogQueryContext)
	if err != nil {
		return err
	}

	// 检查权限
	err = isPermission(b0, reqUser)
	if err != nil {
		return err
	}
	if (req.Status == uint8(entity2.BlogDeleted) || req.Status == uint8(entity2.BlogNotice)) && reqUser.Role < entity2.RoleAdmin {
		return errors.ErrUnauthorized.WithMessage("没有权限将博客封禁或设为公告")
	}

	b1 := blog.NewBlog(blog.WithId(req.Id),
		blog.WithContent(req.Content),
		blog.WithTitle(req.Title),
		blog.WithStatus(entity2.BlogStatus(req.Status)),
		blog.WithProblemId(req.ProblemId))

	return b1.Update()
}
