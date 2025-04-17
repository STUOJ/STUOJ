package collection

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"STUOJ/internal/model/querycontext"
	"sort"
)

// 根据ID更新题单
func Update(req request.UpdateCollectionReq, reqUser model.ReqUser) error {
	queryContext := querycontext.CollectionQueryContext{}
	queryContext.Id.Add(req.ID)
	c0, _, err := collection.Query.SelectOne(queryContext)
	if err != nil {
		return err
	}
	err = isPermission(c0, reqUser)
	if err != nil {
		return err
	}
	c := collection.NewCollection(collection.WithId(uint64(req.ID)),
		collection.WithTitle(req.Title),
		collection.WithDescription(req.Description),
		collection.WithStatus(entity.CollectionStatus(req.Status)),
	)
	return c.Update()
}

// 给题单添加题目
func UpdateProblem(req request.UpdateCollectionProblemReq, reqUser model.ReqUser) error {
	// 查询题单
	queryContext := querycontext.CollectionQueryContext{}
	queryContext.Id.Add(req.CollectionID)
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
		problemIds = append(problemIds, i.ProblemID)
	}

	// 查询题目
	query := querycontext.ProblemQueryContext{}
	query.Id.Set(problemIds)
	problems, _, err := problem.Query.Select(query)
	if err != nil {
		return err
	}
	if len(problems) != len(req.Problem) {
		return errors.ErrUnauthorized.WithMessage("有题目不存在")
	}
	for _, i := range problems {
		if i.Status != entity.ProblemPublic {
			return errors.ErrUnauthorized.WithMessage("有题目不是公开状态")
		}
	}
	return c0.UpdateProblem(problemIds)
}

func UpdateUser(req request.UpdateCollectionUserReq, reqUser model.ReqUser) error {
	// 查询题单
	queryContext := querycontext.CollectionQueryContext{}
	queryContext.Id.Add(req.CollectionID)
	c0, _, err := collection.Query.SelectOne(queryContext)
	if err != nil {
		return err
	}
	if c0.UserId != uint64(reqUser.ID) {
		return errors.ErrUnauthorized.WithMessage("没有权限修改该题单的合作者")
	}
	// 查询用户
	query := querycontext.UserQueryContext{}
	query.Id.Set(req.UserIDS)
	count, err := user.Query.Count(query)
	if err != nil {
		return err
	}
	if count != int64(len(req.UserIDS)) {
		return errors.ErrUnauthorized.WithMessage("有用户不存在")
	}
	return c0.UpdateUser(req.UserIDS)
}
