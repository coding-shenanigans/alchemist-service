package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type ListWishListResponse struct {
	WishLists []*model.WishList `json:"wishLists"`
}
