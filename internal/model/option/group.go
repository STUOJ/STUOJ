package option

type GroupField interface {
	SetField(field string)
	Field() string
	Verify() bool
}
