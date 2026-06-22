package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type CreateItemRequest struct {
	Item *model.Item `json:"item"`
}

// Validates the request fields.
func (r *CreateItemRequest) Validate() error {
	// TODO: Implement validation for the request.
	return nil
}
