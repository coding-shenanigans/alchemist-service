package dto

import (
	"github.com/coding-shenanigans/alchemist-service/internal/validator"
)

type UpdateItemRequest struct {
	Url              *string            `json:"url"`
	Name             *string            `json:"name"`
	Price            *float64           `json:"price"`
	Status           *string            `json:"status"`
	ReservedByUserId NullableField[int] `json:"reservedByUserId"`
}

// Validates the request fields.
func (r *UpdateItemRequest) Validate() error {
	if r.Url != nil {
		if err := validator.ValidateUrl(*r.Url); err != nil {
			return err
		}
	}

	if r.Name != nil {
		if err := validator.ValidateItemName(*r.Name); err != nil {
			return err
		}
	}

	if r.Status != nil {
		if err := validator.ValidateItemStatus(*r.Status); err != nil {
			return err
		}
	}

	return nil
}
