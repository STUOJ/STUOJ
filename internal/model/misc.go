package model

import (
	"STUOJ/internal/db/entity"
	"STUOJ/utils"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Page  uint64 `json:"page,omitempty"`
	Size  uint64 `json:"size,omitempty"`
	Total uint64 `json:"total"`
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

// 时间范围
type Period struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// 检查时间范围
func (p *Period) Check() error {
	if p.StartTime.After(p.EndTime) {
		return errors.New("开始时间不能晚于结束时间")
	}
	return nil
}

// 从字符串解析时间范围
func (p *Period) FromString(startTimeStr string, endTimeStr string, layout string) error {
	var err error

	if startTimeStr == "" || endTimeStr == "" {
		return errors.New("参数错误")
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return errors.New("加载时区失败")
	}

	p.StartTime, err = time.ParseInLocation(layout, startTimeStr, location)
	if err != nil {
		log.Println(err)
		return errors.New("开始时间格式错误")
	}
	p.EndTime, err = time.ParseInLocation(layout, endTimeStr, location)
	if err != nil {
		log.Println(err)
		return errors.New("结束时间格式错误")
	}

	return nil
}

func (p *Period) GetPeriod(c *gin.Context) error {
	var err error

	// 读取参数
	startTimeStr := c.Query("start-time")
	endTimeStr := c.Query("end-time")

	// 解析时间范围
	err = p.FromString(startTimeStr, endTimeStr, utils.DATETIME_LAYOUT)
	if err != nil {
		return err
	}

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return err
	}

	return nil
}

type ReqUser struct {
	ID   int64
	Role entity.Role
}

func (r *ReqUser) Parse(c *gin.Context) {
	role, exist := c.Get("req_user_id")
	if !exist {
		role = entity.RoleVisitor
	}
	id, exist := c.Get("req_user_role")
	if !exist {
		id = 0
	}
	r.ID = id.(int64)
	r.Role = role.(entity.Role)
}

func NewReqUser(c *gin.Context) *ReqUser {
	r := &ReqUser{}
	r.Parse(c)
	return r
}
