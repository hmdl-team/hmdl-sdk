package sdk

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Lấy HandlerContext từ echo Context
func GetHandlerContext(c echo.Context) *HandlerContext {
	return c.(*HandlerContext)
}

type HandlerContext struct {
	echo.Context
}

func (s *HandlerContext) BadRequest(err error) error {
	code := http.StatusBadRequest
	return s.JSON(code, NewErrorResponse(code, err))
}

func (s *HandlerContext) Unauthorized(err error) error {
	code := http.StatusUnauthorized
	return s.JSON(code, NewErrorResponse(code, err))
}

func (s *HandlerContext) NotFound(err error) error {
	code := http.StatusNotFound
	return s.JSON(code, NewErrorResponse(code, err))
}

func (s *HandlerContext) Conflict(err error) error {
	code := http.StatusConflict
	return s.JSON(code, NewErrorResponse(code, err))
}

func (s *HandlerContext) InternalServerError(err error) error {
	code := http.StatusInternalServerError
	return s.JSON(code, NewErrorResponse(code, err))
}

func (s *HandlerContext) Ok(data interface{}) error {
	code := http.StatusOK
	return s.JSON(code, NewSuccessResponse(data))
}

func (s *HandlerContext) Send(response *Response) error {
	return s.JSON(response.StatusCode, response)
}

// Hàm xử lý lỗi chung của hệ thống để trả về cho client.
// Nguyên tắc dựa vào statusCode của appError để xử lý.
// Mặc định mọi error thuần sẽ trả về 500.
func (s *HandlerContext) HandleError(err error) error {
	if err, ok := err.(*AppError); ok {
		return s.JSON(err.StatusCode,
			NewErrorResponse(err.StatusCode, err.rootCause).
				WithMessage(err.Message))
	}

	return s.InternalServerError(err)
}

// Hàm lấy về object user sau khi đăng nhập bằng
func (s *HandlerContext) GetUser() IUser {
	u := s.Get("user")
	user, ok := u.(IUser)
	if !ok {
		return nil
	}

	return user
}

// Hàm lấy về userId sau khi đăng nhập bằng
func (s *HandlerContext) GetUserUid() int {
	user := s.GetUser()
	if user == nil {
		return 0
	}

	return user.GetUid()
}

// Hàm lấy về user role sau khi đăng nhập bằng
func (s *HandlerContext) GetUserRole() string {
	user := s.GetUser()
	if user == nil {
		return ""
	}

	return user.GetRole()
}
