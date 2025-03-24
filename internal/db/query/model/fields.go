package model

type FieldSelector interface {
	SelectedColumns() []string
}
