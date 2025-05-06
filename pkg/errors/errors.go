package errors

import "net/http"

type Error struct {
	HttpStatus int
	Code       int
	Message    string
	Data       any
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(option ...Option) Error {
	err := ErrInternalServer
	for _, o := range option {
		o(&err)
	}
	return err
}

type Option func(*Error)

func WithHttpStatus(status int) Option {
	return func(err *Error) {
		err.HttpStatus = status
	}
}

func WithCode(code int) Option {
	return func(err *Error) {
		err.Code = code
	}
}

func WithMessage(msg string) Option {
	return func(err *Error) {
		err.Message = msg
	}
}

func WithData(data any) Option {
	return func(err *Error) {
		err.Data = data
	}
}

func WithError(err error) Option {
	return func(e *Error) {
		e.Message = err.Error()
	}
}

// WithHttpStatus 返回一个新的Error实例并设置HTTP状态码
// 不会修改原始错误值
func (e *Error) WithHttpStatus(httpStatus int) *Error {
	newErr := *e
	newErr.HttpStatus = httpStatus
	return &newErr
}

// WithCode 返回一个新的Error实例并设置错误码
// 不会修改原始错误值
func (e *Error) WithCode(code int) *Error {
	newErr := *e
	newErr.Code = code
	return &newErr
}

// WithMessage 返回一个新的Error实例并设置错误消息
// 不会修改原始错误值
func (e *Error) WithMessage(msg string) *Error {
	newErr := *e
	newErr.Message = msg
	return &newErr
}

// WithData 返回一个新的Error实例并设置数据
// 不会修改原始错误值
func (e *Error) WithData(data any) *Error {
	newErr := *e
	newErr.Data = data
	return &newErr
}

// WithError 返回一个新的Error实例并设置错误消息
// 不会修改原始错误值
func (e *Error) WithError(err error) *Error {
	newErr := *e
	newErr.Message = err.Error()
	return &newErr
}

// WithErrors 返回一个新的Error实例并追加多个错误消息
// 不会修改原始错误值
func (e *Error) WithErrors(errs []error) *Error {
	newErr := *e
	for _, err := range errs {
		newErr.Message += err.Error() + "\n"
	}
	return &newErr
}

func (e *Error) IsOK() bool {
	return e.Code == 0
}

var (
	ErrInternalServer = Error{
		HttpStatus: http.StatusInternalServerError,
		Code:       1000,
		Message:    "服务器内部错误",
	}
	ErrValidation   = NewError(WithHttpStatus(http.StatusBadRequest), WithCode(1001), WithMessage("请求参数验证失败"))
	ErrUnauthorized = NewError(WithHttpStatus(http.StatusForbidden), WithCode(1002), WithMessage("未授权访问"))
	ErrNotFound     = NewError(WithHttpStatus(http.StatusNotFound), WithCode(1003), WithMessage("资源未找到"))
	ErrId           = ErrValidation.WithMessage("id 不合法")
	ErrStatus       = ErrValidation.WithMessage("status 不合法")
	ErrUserId       = ErrValidation.WithMessage("user_id 不合法")
	ErrProblemId    = ErrValidation.WithMessage("problem_id 不合法")
)
