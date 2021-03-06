package sdk

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	StatusCode    int            `json:"code"`
	Message       string         `json:"message,omitempty"`
	Data          interface{}    `json:"data,omitempty"`
	Input         interface{}    `json:"input,omitempty"`
	InternalError string         `json:"internal_error,omitempty"`
	Tokens        *TokenResponse `json:"tokens,omitempty"`
	Paging        interface{}    `json:"paging,omitempty"`
}
func (s *Response) Success() bool {
	return s.StatusCode <= 399
}

func (s *Response) WithMessage(message string) *Response {
	s.Message = message
	return s
}

func (s *Response) WithInput(input interface{}) *Response {
	s.Input = input
	return s
}

func (s *Response) WithTokens(tokens *TokenResponse) *Response {
	s.Tokens = tokens
	return s
}

func (s *Response) WithError(err error) *Response {
	if err != nil {
		if ae, ok := err.(*AppError); ok {
			if s.Message == "" {
				s.Message = ae.Message
			}
			if ae.rootCause != nil {
				s.InternalError = ae.rootCause.Error()
			}
		} else {
			s.InternalError = err.Error()
		}
	}
	return s
}

func (s *Response) WithPaging(paging interface{}) *Response {
	s.Paging = paging
	return s
}

func (s *Response) WithData(data interface{}) *Response {
	s.Data = data
	return s
}

func (s *Response) WithStatusCode(statusCode int) *Response {
	s.StatusCode = statusCode
	return s
}

func (s *Response) ToBytes() []byte {
	data, _ := json.Marshal(*s)
	return data
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func NewErrorResponse(statusCode int, err error) *Response {
	r := &Response{
		StatusCode: statusCode,
	}
	r.WithError(err)
	return r
}

func ResponseWithCode(c echo.Context, code int, errMsg ...string) error {
	var msg string
	if len(errMsg) == 0 {
		msg = http.StatusText(code)
	} else {
		msg = errMsg[0]
	}
	return c.JSON(code, Response{
		StatusCode: code,
		Message:    msg,
	})
}

func ResponseDataMessage(c echo.Context, mesage string, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    mesage,
		Data:       data,
	})
}

func ResponseData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       data,
	})
}
