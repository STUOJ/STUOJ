package collection

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/errors"
	"STUOJ/pkg/utils"
	"slices"
)

func isPermission(c collection.Collection, reqUser request.ReqUser) error {
	if c.UserId.Value() != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		query := querycontext.CollectionQueryContext{}
		query.Id.Add(c.Id.Value())
		_, map_, err := collection.Query.SelectOne(query, collection.QueryUserId())
		userIds, err := utils.StringToInt64Slice(string(map_["collection_user_id"].([]uint8)))
		if err != nil {
			return err
		}
		if !slices.Contains(userIds, reqUser.Id) {
			return errors.ErrUnauthorized.WithMessage("没有权限操作该题单")
		}
	}
	return nil
}
