package dto

import "github.com/coding-shenanigans/alchemist-service/internal/auth"

type SignupResponse struct {
	UserSession *auth.UserSession `json:"userSession"`
}
