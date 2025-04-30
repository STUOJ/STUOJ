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
	for _, selector := range *s {
		str += selector.String() + ","
	}
	if len(str) == 0 {
		return ""
	}
	return str[:len(str)-1]
}
