package collection

import (
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
	"slices"
)

func isPermission(c collection.Collection, reqUser model.ReqUser) error {
	if c.UserId != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		query := querycontext.CollectionQueryContext{}
		query.Id.Add(c.Id)
		_, map_, err := collection.Query.SelectOne(query, collection.QueryUserId())
		userIds, err := utils.StringToInt64Slice(map_["collection_user_id"].(string))
		if err != nil {
			return err
		}
		if !slices.Contains(userIds, reqUser.Id) {
			return errors.ErrUnauthorized.WithMessage("没有权限操作该题单")
		}
	}
	return nil
}
