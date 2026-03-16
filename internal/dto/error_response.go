package dto

import "github.com/coding-shenanigans/alchemist-service/internal/exception"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error *AppError `json:"error"`
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: &AppError{
			Code:    code,
			Message: message,
		},
	}
}

func NewErrorResponseFromApiError(apiErr *exception.ApiError) *ErrorResponse {
	return &ErrorResponse{
		Error: &AppError{
			Code:    apiErr.Status(),
			Message: apiErr.Error(),
		},
	}
}
