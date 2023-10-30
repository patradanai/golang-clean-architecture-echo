package errs

import (
	"fmt"
)

type Errors interface {
	error
	Error() string
	Status() int
	Causes() interface{}
	Line() string
	DescLine() string
	Code() int
}

type MetaError struct {
	httpCode int
	code     int
	message  string
	line     string
	desc     string
}

func (e *MetaError) Error() string {
	return fmt.Sprintf("%d: %s - %s", e.code, e.desc, e.message)
}

func (e *MetaError) Status() int {
	return e.code
}

// RestError Causes
func (e *MetaError) Causes() interface{} {
	return e.desc
}

func (e *MetaError) Code() int {
	return e.code
}

func (e *MetaError) Line() string {
	return e.line
}

func (e *MetaError) DescLine() string {
	return e.desc + ":" + e.line
}

func (e *MetaError) String() string {
	code := fmt.Sprintf(", Error Code: %d", e.code)
	linecode := fmt.Sprintf(", Line : %s", e.line)

	return fmt.Sprintf("%s%s%s", e.desc, code, linecode)

}

// Wrap error
func WrapError(errorDesc ErrorDesc, params ...string) Errors {
	return &MetaError{
		httpCode: errorDesc.httpCode,
		code:     errorDesc.code,
		desc:     errorDesc.message,
		line:     line(),
	}
}

func WrapDError(err error, code int, params ...string) Errors {
	return &MetaError{
		code: code,
		desc: err.Error(),
		line: line(),
	}
}
