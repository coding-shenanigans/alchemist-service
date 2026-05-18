package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type GetWishListResponse struct {
	WishList *model.WishList `json:"wishList"`
}
