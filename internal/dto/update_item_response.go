package dto

import "github.com/coding-shenanigans/alchemist-service/internal/model"

type UpdateItemResponse struct {
	Item *model.Item `json:"item"`
}
