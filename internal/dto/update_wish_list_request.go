package dto

import (
	"github.com/coding-shenanigans/alchemist-service/internal/validator"
)

type UpdateWishListRequest struct {
	Name       *string `json:"name"`
	Visibility *string `json:"visibility"`
}

// Validates the request fields.
func (r *UpdateWishListRequest) Validate() error {
	if r.Name != nil {
		if err := validator.ValidateWishListName(*r.Name); err != nil {
			return err
		}
	}

	if r.Visibility != nil {
		if err := validator.ValidateWishListVisibility(*r.Visibility); err != nil {
			return err
		}
	}

	return nil
}
