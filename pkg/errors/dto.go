package errs

type (
	ResponseError struct {
		Errors ErrorResponseBody `json:"errors"`
	}

	ErrorResponseBody struct {
		Code     int    `json:"code"`
		Message  string `json:"message"`
		DescLine string `json:"desc_line"`
	}
)
