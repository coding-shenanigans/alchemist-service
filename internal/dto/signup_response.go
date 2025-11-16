package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type SignupResponse struct {
	User *model.User `json:"user"`
}
