package queryfield

import (
	"STUOJ/internal/infrastructure/persistence/entity/field"
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
