package blog

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
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
	if b0.UserId != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	if (req.Status == uint8(entity.BlogBanned) || req.Status == uint8(entity.BlogNotice)) && reqUser.Role < entity.RoleAdmin {
		return errors.ErrUnauthorized.WithMessage("没有权限将博客封禁或设为公告")
	}

	b1 := blog.NewBlog(blog.WithId(req.Id),
		blog.WithContent(req.Content),
		blog.WithTitle(req.Title),
		blog.WithStatus(entity.BlogStatus(req.Status)),
		blog.WithProblemId(req.ProblemId))

	return b1.Update()
}
