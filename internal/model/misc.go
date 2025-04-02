package model

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/query/option"
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

func (f *Field[T]) Set(value interface{}) error {
	if v, ok := value.(T); ok {
		f.exist = true
		f.value = v
		return nil
	}
	return fmt.Errorf("cannot set value of type %T to field of type %T", value, f.value)
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

	f.Set(ptr)

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

func (f *FieldList[T]) Set(value interface{}) error {
	if v, ok := value.([]T); ok {
		f.exist = true
		f.value = v
		return nil
	}
	return fmt.Errorf("cannot set value of type %T to field of type %T", value, f.value)
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

type QueryContext interface {
	Parse(c gin.Context)
	GenerateOptions() *option.QueryOptions
}

type QuerySort struct {
	Order   Field[string]
	OrderBy Field[string]
}

func (s *QuerySort) Parse(c *gin.Context) {
	s.Order.Parse(c, "order")
	s.OrderBy.Parse(c, "order_by")
}

func (s *QuerySort) InsertOptions(options *option.QueryOptions) *option.QueryOptions {
	if s.Order.Exist() && s.OrderBy.Exist() {
		var order option.SortOrder
		if s.Order.Value() == "asc" {
			order = option.OrderAsc
		} else {
			order = option.OrderDesc
		}
		options.Sort = option.NewSortQuery(s.OrderBy.Value(), order)
	}
	return options
}

type QueryPage struct {
	Page  Field[int]
	Size  Field[int]
	Total Field[int]
}

func (p *QueryPage) Parse(c *gin.Context) {
	p.Page.Parse(c, "page")
	p.Size.Parse(c, "size")
	p.Total.Parse(c, "total")
}

func (p *QueryPage) InsertOptions(options *option.QueryOptions) *option.QueryOptions {
	if p.Page.Exist() && p.Size.Exist() {
		options.Page = option.NewPagination(p.Page.Value(), p.Size.Value())
	}
	return options
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
