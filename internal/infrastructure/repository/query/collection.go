package query

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
)

var (
	CollectionAllField    = field.NewCollectionField().SelectAll()
	CollectionSimpleField = field.NewCollectionField().
				SelectId().
				SelectTitle().
				SelectUserId().
				SelectStatus()
	CollectionListItemField = field.NewCollectionField().
				SelectId().
				SelectTitle().
				SelectUserId().
				SelectStatus().
				SelectCreateTime().
				SelectUpdateTime()
)
