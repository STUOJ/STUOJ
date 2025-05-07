package blog

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

// 插入博客
func Insert(req request.CreateBlogReq, reqUser model.ReqUser) (int64, error) {
	blog := blog.NewBlog(blog.WithContent(req.Content),
		blog.WithTitle(req.Title),
		blog.WithUserId(reqUser.Id),
		blog.WithStatus(entity.BlogStatus(req.Status)),
		blog.WithProblemId(req.ProblemId),
	)

	// 检查权限
	if (blog.Status.Value() == entity.BlogDeleted || blog.Status.Value() == entity.BlogNotice) && reqUser.Role < entity.RoleAdmin {
		return 0, errors.ErrUnauthorized.WithMessage("没有权限创建被封禁或公告的博客")
	}

	if blog.ProblemId.Value() != 0 {
		problemQueryContext := querycontext.ProblemQueryContext{}
		problemQueryContext.Id.Add(blog.ProblemId.Value())
		problemQueryContext.Field.SelectStatus()
		problem, _, err := problem.Query.SelectOne(problemQueryContext)
		if err != nil {
			return 0, errors.ErrNotFound.WithMessage("找不到对应的题目")
		}
		if problem.Status.Value() < entity.ProblemPublic {
			return 0, errors.ErrUnauthorized.WithMessage("没有权限创建对应题目未公开的博客")
		}
	}
	return blog.Create()
}
