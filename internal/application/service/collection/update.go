package collection

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/user"
	entity "STUOJ/internal/infrastructure/persistence/entity"
	querycontext2 "STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/errors"
	"sort"
)

// Update 根据Id更新题单
func Update(req request.UpdateCollectionReq, reqUser request.ReqUser) error {
	queryContext := querycontext2.CollectionQueryContext{}
	queryContext.Id.Add(req.Id)
	c0, _, err := collection.Query.SelectOne(queryContext)
	if err != nil {
		return err
	}
	err = isPermission(c0, reqUser)
	if err != nil {
		return err
	}
	c := collection.NewCollection(collection.WithId(int64(req.Id)),
		collection.WithTitle(req.Title),
		collection.WithDescription(req.Description),
		collection.WithStatus(entity.CollectionStatus(req.Status)),
	)
	return c.Update()
}

// UpdateProblem 给题单添加题目
func UpdateProblem(req request.UpdateCollectionProblemReq, reqUser request.ReqUser) error {
	// 查询题单
	queryContext := querycontext2.CollectionQueryContext{}
	queryContext.Id.Add(int64(req.CollectionId))
	c0, _, err := collection.Query.SelectOne(queryContext)
	if err != nil {
		return err
	}
	err = isPermission(c0, reqUser)
	if err != nil {
		return err
	}

	// 按serial升序排序
	sort.Slice(req.Problem, func(i, j int) bool {
		return req.Problem[i].Serial < req.Problem[j].Serial
	})

	problemIds := make([]int64, len(req.Problem))
	for _, i := range req.Problem {
		problemIds = append(problemIds, int64(i.ProblemId))
	}

	// 查询题目
	query := querycontext2.ProblemQueryContext{}
	query.Id.Set(problemIds)
	problems, _, err := problem.Query.Select(query)
	if err != nil {
		return err
	}
	if len(problems) != len(req.Problem) {
		return errors.ErrUnauthorized.WithMessage("有题目不存在")
	}
	for _, i := range problems {
		if i.Status.Value() != entity.ProblemPublic {
			return errors.ErrUnauthorized.WithMessage("有题目不是公开状态")
		}
	}
	return c0.UpdateProblem(problemIds)
}

func UpdateUser(req request.UpdateCollectionUserReq, reqUser request.ReqUser) error {
	// 查询题单
	queryContext := querycontext2.CollectionQueryContext{}
	queryContext.Id.Add(int64(req.CollectionId))
	c0, _, err := collection.Query.SelectOne(queryContext)
	if err != nil {
		return err
	}
	if c0.UserId.Value() != int64(reqUser.Id) {
		return errors.ErrUnauthorized.WithMessage("没有权限修改该题单的合作者")
	}
	// 查询用户
	query := querycontext2.UserQueryContext{}
	// 将int64切片转换为uint64切片
	userIds := make([]int64, len(req.UserIds))
	for i, id := range req.UserIds {
		userIds[i] = id
	}
	query.Id.Add(userIds...)
	count, err := user.Query.Count(query)
	if err != nil {
		return err
	}
	if count != int64(len(req.UserIds)) {
		return errors.ErrUnauthorized.WithMessage("有用户不存在")
	}
	return c0.UpdateUser(userIds)
}
