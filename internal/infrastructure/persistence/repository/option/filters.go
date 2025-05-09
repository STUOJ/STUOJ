package option

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type FilterOperator string

const (
	OpEqual     FilterOperator = "="
	OpGreater   FilterOperator = ">"
	OpGreaterEq FilterOperator = ">="
	OpLess      FilterOperator = "<"
	OpLessEq    FilterOperator = "<="
	OpIn        FilterOperator = "IN"
	OpNotIn     FilterOperator = "NOT IN"
	OpLike      FilterOperator = "LIKE"
	OpNotLike   FilterOperator = "NOT LIKE"
	OpIsNull    FilterOperator = "IS NULL"
	OpIsNotNull FilterOperator = "IS NOT NULL"
	OpExtra     FilterOperator = ""
)

type FilterCondition struct {
	Field    string
	Operator FilterOperator
	Value    any
}

func (f *FilterCondition) String() string {
	return fmt.Sprintf("FilterCondition{Field: %s, Operator: %s, Value: %v}", f.Field, f.Operator, f.Value)
}

type Filters struct {
	Conditions []FilterCondition
	Errors     []error
}

func NewFilters() *Filters { return &Filters{} }

func (f *Filters) String() string {
	return fmt.Sprintf("Filters{Conditions: %v, Errors: %v}", f.Conditions, f.Errors)
}

func (f *Filters) Add(field string, operator FilterOperator, values ...any) (err error) {
	if field == "" {
		err = fmt.Errorf("field cannot be empty")
		return
	}
	var value any
	switch operator {
	case OpIsNull, OpIsNotNull:
		if len(values) > 0 {
			err = fmt.Errorf("%s operator does not require value", operator)
			return
		}
	case OpIn, OpNotIn:
		if len(values) == 0 {
			err = fmt.Errorf("%s operator requires at least one value", operator)
			return
		}
		if reflect.TypeOf(values[0]).Kind() != reflect.Slice {
			value = values
		} else {
			value = values[0]
		}
	case OpLike, OpNotLike:
		if len(values) == 0 {
			return fmt.Errorf("%s operator requires a value", operator)
		}
		if _, ok := values[0].(string); !ok {
			return fmt.Errorf("%s requires string value (got %T)", operator, values[0])
		}
	case OpExtra:
		value = values
	default:
		if len(values) == 0 {
			err = fmt.Errorf("%s operator requires a value", operator)
			return
		}
		value = values[0]
	}

	f.Conditions = append(f.Conditions, FilterCondition{
		Field:    field,
		Operator: operator,
		Value:    value,
	})
	return
}

func (f *Filters) AddFiter(filter ...FilterCondition) {
	f.Conditions = append(f.Conditions, filter...)
}

func (f *Filters) GenerateWhere() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, condition := range f.Conditions {
			db = f.applyCondition(db, condition)
		}
		return db
	}
}

func (f *Filters) applyCondition(db *gorm.DB, c FilterCondition) *gorm.DB {
	switch c.Operator {

	case OpEqual, OpGreater, OpGreaterEq, OpLess, OpLessEq:
		return db.Where(fmt.Sprintf("%s %s ?", c.Field, c.Operator), c.Value)
	case OpIn, OpNotIn:
		if reflect.TypeOf(c.Value).Kind() != reflect.Slice {
			f.Errors = append(f.Errors,
				fmt.Errorf("IN操作需要slice类型参数(字段:%s)", c.Field))
			return db
		}
		return db.Where(fmt.Sprintf("%s %s ?", c.Field, c.Operator), c.Value)
	case OpLike, OpNotLike:
		if _, ok := c.Value.(string); !ok {
			f.Errors = append(f.Errors, fmt.Errorf("%s operator requires string value", c.Operator))
		}
		return db.Where(fmt.Sprintf("%s %s ?", c.Field, c.Operator), "%"+c.Value.(string)+"%")
	case OpIsNull, OpIsNotNull:
		return db.Where(fmt.Sprintf("%s %s", c.Field, c.Operator))
	case OpExtra:
		// 计算Field中问号的数量
		questionMarkCount := strings.Count(c.Field, "?")

		// 如果没有问号，直接传递所有参数
		if questionMarkCount == 0 {
			return db.Where(c.Field, c.Value)
		}

		// 如果只有一个问号，直接使用参数
		if questionMarkCount == 1 {
			return db.Where(c.Field, c.Value)
		}

		// 如果有多个问号，需要确保参数是切片且长度匹配
		values, ok := c.Value.([]any)
		if !ok || len(values) != questionMarkCount {
			f.Errors = append(f.Errors,
				fmt.Errorf("问号数量(%d)与参数数量(%d)不匹配(字段:%s)",
					questionMarkCount, len(values), c.Field))
			return db
		}

		// 问号数量与值数量匹配，直接使用Where
		return db.Where(c.Field, values...)
	default:
		f.Errors = append(f.Errors,
			fmt.Errorf("不支持的运算符:%s (字段:%s)", c.Operator, c.Field))
	}
	return db
}
