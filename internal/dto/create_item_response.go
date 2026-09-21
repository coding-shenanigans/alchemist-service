package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type CreateItemResponse struct {
	Item *model.ItemWithUser `json:"item"`
}
