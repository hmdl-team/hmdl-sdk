package sdk

import (
	"github.com/dgrijalva/jwt-go"
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

// Hàm lấy uid từ access_token và đã được check exist trong database,
// nếu handler được bọc middleware authentication
func (s *HandlerContext) GetUid() *int {
	tokenData := s.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*JwtClaims)

	if claims !=nil {
		return &claims.UserId
	}


	return nil
}
