package sdkcm

import "net/http"

type Response struct {
	StatusCode    int            `json:"status_code"`
	Data          interface{}    `json:"data,omitempty"`
	Message       string         `json:"message,omitempty"`
	Input         interface{}    `json:"input,omitempty"`
	InternalError string         `json:"internal_error,omitempty"`
	Tokens        *TokenResponse `json:"tokens,omitempty"`
	Paging        interface{}    `json:"paging,omitempty"`
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
