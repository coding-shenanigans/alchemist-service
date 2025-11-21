package exception

type ApiError struct {
	status  int
	message string
}

func NewApiError(status int, message string) *ApiError {
	return &ApiError{
		status:  status,
		message: message,
	}
}

func (e *ApiError) Status() int {
	return e.status
}

func (e *ApiError) Error() string {
	return e.message
}
