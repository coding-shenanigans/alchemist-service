package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type UpdateWishListResponse struct {
	WishList *model.WishList `json:"wishList"`
}
