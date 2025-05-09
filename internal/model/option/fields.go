package option

type FieldSelector interface {
	SelectedColumns() []string
	AddSelect(selector ...Selector)
}
