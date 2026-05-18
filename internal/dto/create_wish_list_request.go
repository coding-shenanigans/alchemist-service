package dto

import (
	"github.com/coding-shenanigans/alchemist-service/internal/model"
	"github.com/coding-shenanigans/alchemist-service/internal/validator"
)

type CreateWishListRequest struct {
	WishList *model.WishList `json:"wishList"`
}

// Validates the request fields.
func (r *CreateWishListRequest) Validate() error {
	if err := validator.ValidateWishListName(r.WishList.Name); err != nil {
		return err
	}

	err := validator.ValidateWishListVisibility(r.WishList.Visibility)
	if err != nil {
		return err
	}

	return nil
}
