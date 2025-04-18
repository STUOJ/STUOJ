package collection

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"

	"STUOJ/internal/domain/collection"

	"STUOJ/internal/model"
)

// 插入题单
func Insert(req request.CreateCollectionReq, reqUser model.ReqUser) (uint64, error) {
	c := collection.NewCollection(collection.WithTitle(req.Title), collection.WithDescription(req.Description), collection.WithUserId(uint64(reqUser.Id)), collection.WithStatus(entity.CollectionStatus(req.Status)))
	return c.Create()
}
