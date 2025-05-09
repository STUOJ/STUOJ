package option

import "fmt"

type Selector struct {
	Query string
	Arg   []any
}

func (s *Selector) SelectedColumns() []any {
	return s.Arg
}

func (s *Selector) AddArg(arg ...any) {
	s.Arg = append(s.Arg, arg...)
}

func (s *Selector) SetQuery(query string) {
	s.Query = query
}

func (s *Selector) String() string {
	return fmt.Sprintf(s.Query, s.Arg...)
}

func NewSelector(query string, arg ...any) *Selector {
	return &Selector{
		Query: query,
		Arg:   arg,
	}
}

type Selectors []Selector

func (s *Selectors) AddSelector(selector ...Selector) {
	*s = append(*s, selector...)
}

func (s *Selectors) String() string {
	var str string
	for i, selector := range *s {
		if i > 0 {
			str += ","
		}
		str += selector.String()
	}
	return str
}

func (s *Selectors) SelectedColumns() []string {
	var columns []string
	for _, selector := range *s {
		columns = append(columns, selector.String())
	}
	return columns
}
