package blog

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/query"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
	"slices"
)

type BlogPage struct {
	Blogs []response.BlogData `json:"blogs"`
	model.Page
}

// SelectById 根据Id查询博客
func SelectById(id int64, reqUser model.ReqUser) (response.BlogData, error) {
	var resp response.BlogData
	blogQuery := querycontext.BlogQueryContext{}
	blogQuery.Id.Add(id)
	blogQuery.Field = *query.BlogAllField

	domainBlog, _, err := blog.Query.SelectOne(blogQuery)
	if err != nil {
		return resp, err
	}
	resp = domain2Resp(domainBlog)
	userQuery := querycontext.UserQueryContext{}
	userQuery.Id.Add(domainBlog.UserId.Value())
	userQuery.Field = *query.UserSimpleField
	domainUser, _, err := user.Query.SelectOne(userQuery)
	if err == nil {
		resp.User = response.Domain2UserSimpleData(domainUser)
	}
	problemQuery := querycontext.ProblemQueryContext{}
	problemQuery.Id.Add(domainBlog.ProblemId.Value())
	problemQuery.Field = *query.ProblemSimpleField
	_, map_, err := problem.Query.SelectOne(problemQuery, problem.QueryMaxScore(reqUser.Id), problem.QueryTag())
	if err == nil {
		resp.Problem.ProblemSimpleData = response.Map2ProblemSimpleData(map_)
		resp.Problem.ProblemUserScore = response.Map2ProblemUserScore(map_)
	}
	return resp, nil
}

func Select(params request.QueryBlogParams, reqUser model.ReqUser) (BlogPage, error) {
	var resp BlogPage
	query_ := params2Query(params)
	if !query_.Status.Exist() {
		query_.Status.Set([]uint8{uint8(entity.BlogPublic)})
	}
	if (slices.Contains(query_.Status.Value(), uint8(entity.BlogDeleted)) || slices.Contains(query_.Status.Value(), uint8(entity.BlogDraft))) && reqUser.Role < entity.RoleAdmin {
		query_.UserId.Set([]int64{reqUser.Id})
	}
	query_.Field = *query.BlogAllField
	blogs, _, err := blog.Query.Select(query_)
	if err != nil {
		return BlogPage{}, err
	}
	problemIds := make([]int64, len(blogs))
	for i, blog_ := range blogs {
		problemIds[i] = blog_.ProblemId.Value()
	}
	userIds := make([]int64, len(blogs))
	for i, blog_ := range blogs {
		userIds[i] = blog_.UserId.Value()
	}
	problemQueryContext := querycontext.ProblemQueryContext{}
	problemQueryContext.Id.Add(problemIds...)
	problemQueryContext.Field = *query.ProblemSimpleField
	_, problemMap, err := problem.Query.SelectByIds(problemQueryContext, problem.QueryMaxScore(reqUser.Id), problem.QueryTag())

	userQueryContext := querycontext.UserQueryContext{}
	userQueryContext.Id.Add(userIds...)
	userQueryContext.Field = *query.UserSimpleField
	users, _, err := user.Query.Select(userQueryContext)

	for _, blog_ := range blogs {
		var resBlog response.BlogData
		resBlog = domain2Resp(blog_)
		if blog_.ProblemId.Value() != 0 {
			resBlog.Problem.ProblemSimpleData = response.Map2ProblemSimpleData(problemMap[blog_.ProblemId.Value()])
			resBlog.Problem.ProblemUserScore = response.Map2ProblemUserScore(problemMap[blog_.ProblemId.Value()])
		}
		if blog_.UserId.Value() != 0 {
			resBlog.User = response.Domain2UserSimpleData(users[blog_.UserId.Value()])
		}
		resp.Blogs = append(resp.Blogs, resBlog)
	}
	resp.Page.Page = query_.Page.Page
	resp.Page.Size = query_.Page.PageSize
	resp.Page.Total, err = Count(params)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
