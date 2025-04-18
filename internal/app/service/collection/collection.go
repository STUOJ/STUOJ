package collection

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"slices"
)

func isPermission(c collection.Collection, reqUser model.ReqUser) error {
	if c.UserId != uint64(reqUser.Id) && reqUser.Role < entity.RoleAdmin {
		query := querycontext.CollectionQueryContext{}
		query.Id.Add(int64(c.Id))
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
