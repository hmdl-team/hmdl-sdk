package sdk

import (
	"errors"
	"net/http"
)

var (
	ErrPhoneEmpty  = errors.New("Số điện thoại không được rỗng")
	ErrPhoneFormat = errors.New("Số điện thoại không hợp lệ")
	ErrOidEmpty    = errors.New("Không tìm thấy Oid")
)

var (
	ErrDataNotFound = func(err error) *AppError {
		return NewAppError(err).
			WithMessage("Không có dữ liệu").
			WithStatusCode(http.StatusNotFound)
	}
	ErrUnauthorized = func(err error) *AppError {
		return NewAppError(err).
			WithMessage("Bạn chưa đăng nhập").
			WithStatusCode(http.StatusUnauthorized)
	}
)

type AppError struct {
	rootCause  error  `json:"-"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewAppError(rootCause error) *AppError {
	return &AppError{rootCause: rootCause}
}

func (s *AppError) WithMessage(message string) *AppError {
	s.Message = message
	return s
}

func (s *AppError) WithStatusCode(statusCode int) *AppError {
	s.StatusCode = statusCode
	return s
}

func (s *AppError) Error() string {
	return s.Message
}

func NewBadRequestErr(rootCause error) *AppError {
	return NewAppError(rootCause).WithStatusCode(http.StatusBadRequest)
}

func NewNotFoundErr(rootCause error) *AppError {
	return NewAppError(rootCause).WithStatusCode(http.StatusNotFound)
}

func NewConflictErr(rootCause error) *AppError {
	return NewAppError(rootCause).WithStatusCode(http.StatusConflict)
}

func NewUnauthorizedErr(rootCause error) *AppError {
	return NewAppError(rootCause).WithStatusCode(http.StatusUnauthorized).WithMessage("Thông tin đăng nhập không chính xác")
}

func IsNotFoundError(err error) bool {
	if e, ok := err.(*AppError); ok {
		if e.StatusCode == http.StatusNotFound {
			return true
		}
	}
	return false
}
