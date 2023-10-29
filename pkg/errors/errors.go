package errs

type ResponseError struct {
}

type Errors struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	error
}
