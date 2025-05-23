package blog

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/blog"
	entity "STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/errors"
)

func Update(req request.UpdateBlogReq, reqUser request.ReqUser) error {
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
	if (req.Status == uint8(entity.BlogDeleted) || req.Status == uint8(entity.BlogNotice)) && reqUser.Role < entity.RoleAdmin {
		return errors.ErrUnauthorized.WithMessage("没有权限将博客封禁或设为公告")
	}

	b1 := blog.NewBlog(blog.WithId(req.Id),
		blog.WithContent(req.Content),
		blog.WithTitle(req.Title),
		blog.WithStatus(entity.BlogStatus(req.Status)),
		blog.WithProblemId(req.ProblemId))

	return b1.Update()
}
