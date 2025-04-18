package blog

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/model"
	"slices"
)

type BlogPage struct {
	Blogs []response.BlogData `json:"blogs"`
	model.Page
}

// 根据Id查询博客
func SelectById(id uint64, reqUser model.ReqUser) (response.BlogData, error) {
	var res response.BlogData
	blogQuery := querycontext.BlogQueryContext{}
	blogQuery.Id.Add(int64(id))
	blogQuery.Field = *query.BlogAllField

	domainBlog, _, err := blog.Query.SelectOne(blogQuery)
	if err != nil {
		return res, err
	}
	res = domain2response(domainBlog)
	userQuery := querycontext.UserQueryContext{}
	userQuery.Id.Add(int64(domainBlog.UserId))
	userQuery.Field = *query.UserSimpleField
	domainUser, _, err := user.Query.SelectOne(userQuery)
	if err == nil {
		res.User = response.Domain2UserSimpleData(domainUser)
	}
	problemQuery := querycontext.ProblemQueryContext{}
	problemQuery.Id.Add(int64(domainBlog.ProblemId))
	problemQuery.Field = *query.ProblemSimpleField
	_, map_, err := problem.Query.SelectOne(problemQuery, problem.QueryMaxScore(reqUser.Id), problem.QueryTag())
	if err == nil {
		res.Problem.ProblemSimpleData = response.Map2ProblemSimpleData(map_)
		res.Problem.ProblemUserScore = response.Map2ProblemUserScore(map_)
	}
	return res, nil
}

func Select(params request.QueryBlogParams, resUser model.ReqUser) (BlogPage, error) {
	var res BlogPage
	query_ := params2Model(params)
	if !query_.Status.Exist() {
		query_.Status.Set([]int64{int64(entity.BlogPublic)})
	}
	if (slices.Contains(query_.Status.Value(), int64(entity.BlogBanned)) || slices.Contains(query_.Status.Value(), int64(entity.BlogDraft))) && resUser.Role < entity.RoleAdmin {
		query_.UserId.Set([]int64{resUser.Id})
	}
	query_.Field = *query.BlogAllField
	blogs, _, err := blog.Query.Select(query_)
	if err != nil {
		return BlogPage{}, err
	}
	problemIds := make([]int64, len(blogs))
	for i, blog_ := range blogs {
		problemIds[i] = int64(blog_.ProblemId)
	}
	userIds := make([]int64, len(blogs))
	for i, blog_ := range blogs {
		userIds[i] = int64(blog_.UserId)
	}
	problemQueryContext := querycontext.ProblemQueryContext{}
	problemQueryContext.Id.Add(problemIds...)
	problemQueryContext.Field = *query.ProblemSimpleField
	_, problemMap, err := problem.Query.SelectByIds(problemQueryContext, problem.QueryMaxScore(resUser.Id), problem.QueryTag())

	userQueryContext := querycontext.UserQueryContext{}
	userQueryContext.Id.Add(userIds...)
	userQueryContext.Field = *query.UserSimpleField
	users, _, err := user.Query.Select(userQueryContext)

	for _, blog_ := range blogs {
		var resBlog response.BlogData
		resBlog = domain2response(blog_)
		if blog_.ProblemId != 0 {
			resBlog.Problem.ProblemSimpleData = response.Map2ProblemSimpleData(problemMap[int64(blog_.ProblemId)])
			resBlog.Problem.ProblemUserScore = response.Map2ProblemUserScore(problemMap[int64(blog_.ProblemId)])
		}
		if blog_.UserId != 0 {
			resBlog.User = response.Domain2UserSimpleData(users[int64(blog_.UserId)])
		}
		res.Blogs = append(res.Blogs, resBlog)
	}
	res.Page.Page = uint64(query_.Page.Page)
	res.Page.Size = uint64(query_.Page.PageSize)
	total, _ := GetStatistics(params)
	res.Page.Total = uint64(total)
	return res, nil
}
