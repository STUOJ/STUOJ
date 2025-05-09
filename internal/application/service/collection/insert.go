package collection

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
)

// Insert 插入题单
func Insert(req request.CreateCollectionReq, reqUser model.ReqUser) (int64, error) {
	c := collection.NewCollection(collection.WithTitle(req.Title), collection.WithDescription(req.Description), collection.WithUserId(reqUser.Id), collection.WithStatus(entity.CollectionStatus(req.Status)))
	return c.Create()
}
