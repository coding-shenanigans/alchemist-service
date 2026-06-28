package dto

import (
	"github.com/coding-shenanigans/alchemist-service/internal/model"
	"github.com/coding-shenanigans/alchemist-service/internal/validator"
)

type CreateItemRequest struct {
	Item *model.Item `json:"item"`
}

// Validates the request fields.
func (r *CreateItemRequest) Validate() error {
	if err := validator.ValidateUrl(r.Item.Url); err != nil {
		return err
	}

	if err := validator.ValidateItemName(r.Item.Name); err != nil {
		return err
	}

	return nil
}
