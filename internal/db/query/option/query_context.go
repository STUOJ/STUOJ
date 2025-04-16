package option

type QueryContext interface {
	GenerateOptions() *QueryOptions
	GetField() FieldSelector
}

type QueryContextOption func(QueryContext) QueryContext
