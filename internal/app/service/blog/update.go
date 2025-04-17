package blog

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"STUOJ/internal/model/querycontext"
)

func Update(req request.UpdateBlogReq, reqUser model.ReqUser) error {
	blogQueryContext := querycontext.BlogQueryContext{}
	blogQueryContext.Id.Add(int64(req.ID))
	blogQueryContext.Field.SelectId().SelectStatus().SelectUserId()
	blog0, _, err := blog.Query.SelectOne(blogQueryContext)
	if err != nil {
		return err
	}
	err = isPermission(blog0, reqUser)
	if err != nil {
		return err
	}
	if (req.Status == int64(entity.BlogBanned) || req.Status == int64(entity.BlogNotice)) && reqUser.Role < entity.RoleAdmin {
		return errors.ErrUnauthorized.WithMessage("没有权限将博客封禁或设为公告")
	}

	blog_ := blog.NewBlog(blog.WithID(uint64(req.ID)),
		blog.WithContent(req.Content),
		blog.WithTitle(req.Title),
		blog.WithStatus(entity.BlogStatus(req.Status)),
		blog.WithProblemID(uint64(req.ProblemID)))

	return blog_.Update()
}
