package option

type FieldSelector interface {
	SelectedColumns() []string
}
