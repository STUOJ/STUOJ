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

func (e *Error) WithHttpStatus(httpStatus int) *Error {
	e.HttpStatus = httpStatus
	return e
}

func (e *Error) WithCode(code int) *Error {
	e.Code = code
	return e
}

func (e *Error) WithMessage(msg string) *Error {
	e.Message = msg
	return e
}

func (e *Error) WithData(data any) *Error {
	e.Data = data
	return e
}

func (e *Error) WithError(err error) *Error {
	e.Message = err.Error()
	return e
}

func (e *Error) WithErrors(errs []error) *Error {
	for _, err := range errs {
		e.Message += err.Error() + "\n"
	}
	return e
}

func (e *Error) IsOK() bool {
	return e.Code == 0
}

var (
	NoError           = Error{HttpStatus: http.StatusOK, Code: 0, Message: ""}
	ErrInternalServer = Error{
		HttpStatus: http.StatusInternalServerError,
		Code:       1000,
		Message:    "服务器内部错误",
	}
	ErrValidation   = NewError(WithHttpStatus(http.StatusBadRequest), WithCode(1001), WithMessage("请求参数验证失败"))
	ErrUnauthorized = NewError(WithHttpStatus(http.StatusForbidden), WithCode(1002), WithMessage("未授权访问"))
	ErrNotFound     = NewError(WithHttpStatus(http.StatusNotFound), WithCode(1003), WithMessage("资源未找到"))
)
