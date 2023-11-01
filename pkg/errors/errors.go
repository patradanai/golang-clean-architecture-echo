package errs

import (
	"fmt"
	"strings"
)

type Errors interface {
	error

	HttpCode() int

	Error() string

	Status() int

	Stack() interface{}

	Cause() error

	Line() string

	DescLine() string

	Message() string

	Code() int

	FieldError() []FieldError
}

type MetaError struct {
	httpCode int
	code     int
	message  string
	line     string
	causeBy  error
	stack    interface{}
	fields   []FieldError
}

func (e *MetaError) HttpCode() int {
	return e.httpCode
}

func (e *MetaError) Error() string {
	return e.String()
}

func (e *MetaError) Status() int {
	return e.code
}

func (e *MetaError) Message() string {
	return e.message
}

// RestError Causes
func (e *MetaError) Stack() interface{} {
	return e.stack
}

func (e *MetaError) Code() int {
	return e.code
}

func (e *MetaError) Line() string {
	return e.line
}

func (e *MetaError) Cause() error {
	return e.causeBy
}

func (e *MetaError) DescLine() string {
	return e.message + " : " + e.line
}

func (e *MetaError) String() string {
	causeBy := ""
	if e.causeBy != nil {
		causeBy = fmt.Sprintf(", Caused by: %s", e.causeBy.Error())
	}
	code := fmt.Sprintf(", Error Code: %d", e.code)
	linecode := fmt.Sprintf(", Line : %s", e.line)

	return fmt.Sprintf("%s%s%s%s", e.message, code, linecode, causeBy)
}

func (e *MetaError) FieldError() []FieldError {
	return e.fields
}

// Wrap error with errorDesc
func WrapError(errorDesc ErrorDesc, params ...interface{}) Errors {
	return &MetaError{
		httpCode: errorDesc.httpCode,
		code:     errorDesc.code,
		message:  replaceDynamicMessage(errorDesc.message, params...),
		line:     line(),
		causeBy:  nil,
		stack:    nil,
	}
}

// Wrap Error with general error and errorDesc
func WrapDError(err error, errorDesc ErrorDesc, params ...interface{}) Errors {
	return &MetaError{
		httpCode: errorDesc.httpCode,
		code:     errorDesc.code,
		message:  replaceDynamicMessage(errorDesc.message, params...),
		causeBy:  err,
		stack:    nil,
		line:     line(),
	}
}

// Wrap error for Custom Validate and errorDesc
func WrapFieldError(errorDesc ErrorDesc, fields []FieldError, params ...interface{}) Errors {
	return &MetaError{
		httpCode: errorDesc.httpCode,
		code:     errorDesc.code,
		message:  replaceDynamicMessage(errorDesc.message, params...),
		causeBy:  nil,
		stack:    nil,
		line:     line(),
		fields:   fields,
	}
}

func replaceDynamicMessage(plainText string, params ...interface{}) string {
	if strings.Contains(plainText, "%v") {
		return fmt.Sprintf(plainText, params...)
	}
	return fmt.Sprintf(plainText)
}
