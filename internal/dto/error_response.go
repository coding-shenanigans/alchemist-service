package dto

import "github.com/coding-shenanigans/alchemist-service/internal/exception"

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	ErrorInfo *ErrorInfo `json:"error"`
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		ErrorInfo: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	}
}

func NewErrorResponseFromApiError(apiErr *exception.ApiError) *ErrorResponse {
	return &ErrorResponse{
		ErrorInfo: &ErrorInfo{
			Code:    apiErr.Status(),
			Message: apiErr.Error(),
		},
	}
}
