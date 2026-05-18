package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type CreateWishListResponse struct {
	WishList *model.WishList `json:"wishList"`
}
