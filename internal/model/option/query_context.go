package option

type QueryContext interface {
	GenerateOptions() *QueryOptions
	GetField() FieldSelector
	GetExtraFilters() *Filters
}

type QueryContextOption func(QueryContext) QueryContext

type QueryParams struct {
	GroupBy      GroupField
	ExtraFilters Filters
	Page         Pagination
	Sort         Sort
}
