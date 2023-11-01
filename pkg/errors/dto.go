package errs

type (
	ResponseError struct {
		Errors ErrorResponseBody `json:"errors"`
	}

	FieldError struct {
		FieldName   string `json:"field_name"`
		Description string `json:"description"`
	}

	ErrorResponseBody struct {
		Code     int          `json:"code"`
		Message  string       `json:"message"`
		DescLine string       `json:"desc_line"`
		Fields   []FieldError `json:"fields,omitempty"`
	}
)
