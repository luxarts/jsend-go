package jsend

import "fmt"

const (
	StatusSuccess = "success"
	StatusError = "error"
	StatusFail = "fail"
)

// Body JSend specification
// Status: success, fail, error
// Message: Only exists if status == error
type Body struct {
	Status string `json:"status"`
	Message *string `json:"message,omitempty"`
	Data interface{} `json:"data"`
	Code *int `json:"code,omitempty"`
}

func NewError(message string, err error, code ...int) *Body {
	b := Body{
		Status:  StatusError,
		Message: &message,
		Data:    err.Error(),
	}

	if len(code) > 0 {
		b.Code = &(code[0])
	}

	return &b
}

func NewSuccess(data interface{}) *Body {
	return &Body{
		Status: StatusSuccess,
		Data: data,
	}
}

func NewFail(data interface{}) *Body{
	return &Body{
		Status: StatusFail,
		Data: data,
	}
}

func (e *Body) Error() string {
	if e.Status == StatusError{
		if e.Code != nil {
			return fmt.Sprintf("%s (%d)", *e.Message, *e.Code)
		}
		return *e.Message
	} else if e.Status == StatusFail {
		return fmt.Sprintf("%v", e.Data)
	} else {
		return ""
	}
}