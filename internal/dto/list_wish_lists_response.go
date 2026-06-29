package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type ListWishListsResponse struct {
	WishLists []*model.WishList `json:"wishLists"`
}
