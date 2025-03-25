package query

import "STUOJ/internal/db/entity/field"

var (
	CollectionAllField    = field.NewCollectionField().SelectAll()
	CollectionSimpleField = field.NewCollectionField().
				SelectId().
				SelectTitle().
				SelectUserId().
				SelectStatus().
				SelectCreateTime().
				SelectUpdateTime()
)
