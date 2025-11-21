package dto

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func NewErrorResponse(errors ...string) *ErrorResponse {
	return &ErrorResponse{
		Errors: errors,
	}
}
