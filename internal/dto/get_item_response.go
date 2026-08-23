package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type GetItemResponse struct {
	Item *model.ItemWithUser `json:"item"`
}
