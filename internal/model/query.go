package model

import (
	"STUOJ/pkg/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Page  int64 `json:"page,omitempty"`
	Size  int64 `json:"size,omitempty"`
	Total int64 `json:"total"`
}

type Field[T any] struct {
	exist bool
	value T
}

func (f *Field[T]) Exist() bool {
	return f.exist
}

func (f *Field[T]) Value() T {
	return f.value
}

func (f *Field[T]) Set(value T) {
	f.exist = true
	f.value = value
}

func (f *Field[T]) Parse(c *gin.Context, name string) error {
	query := c.Query(name)
	if query == "" {
		return nil
	}

	var tmp T
	var ptr interface{} = tmp

	if err := utils.ConvertStringToType[T](query, &ptr); err != nil {
		return fmt.Errorf("failed to parse value for field %s: %w", name, err)
	}

	if v, ok := ptr.(T); ok {
		f.Set(v)
	} else {
		return fmt.Errorf("type assertion failed: expected %T, got %T", *new(T), ptr)
	}

	f.exist = true
	return nil
}

type FieldList[T any] struct {
	exist bool
	value []T
}

func (f *FieldList[T]) Exist() bool {
	return f.exist
}

func (f *FieldList[T]) Value() []T {
	return f.value
}

func (f *FieldList[T]) Set(value []T) {
	f.exist = true
	f.value = value
}

func (f *FieldList[T]) Add(value ...T) {
	f.exist = true
	f.value = append(f.value, value...)
}

func (f *FieldList[T]) Parse(c *gin.Context, name string) error {
	query := c.Query(name)
	if query == "" {
		return nil
	}
	splQuerys := strings.Split(query, ",")

	var tmp []T

	for _, splQuery := range splQuerys {
		var tmpT T
		var ptr interface{} = tmpT
		if err := utils.ConvertStringToType[T](splQuery, &ptr); err != nil {
			return fmt.Errorf("failed to parse value for field %s: %w", name, err)
		}
		if v, ok := ptr.(T); ok {
			tmp = append(tmp, v)
		} else {
			return fmt.Errorf("type assertion failed: expected %T, got %T", tmpT, ptr)
		}
	}
	f.Set(tmp)
	f.exist = true
	return nil
}
