package service

import (
	"net/http"

	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
	"github.com/coding-shenanigans/alchemist-service/internal/repository"
)

type WishListService struct {
	wishListRepository *repository.WishListRepository
}

func NewWishListService(
	wishListRepository *repository.WishListRepository,
) *WishListService {
	return &WishListService{
		wishListRepository: wishListRepository,
	}
}

func (s *WishListService) CreateWishList(
	wishList *model.WishList,
) (*model.WishList, *exception.ApiError) {
	// TODO: Implement function.
	return nil, exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}

func (s *WishListService) GetWishList(
	id int,
) (*model.WishList, *exception.ApiError) {
	// TODO: Implement function.
	return nil, exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}

func (s *WishListService) ListWishLists(
	pageSize int, pageToken string, filter string,
) ([]*model.WishList, *exception.ApiError) {
	// TODO: Implement function.
	return nil, exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}

func (s *WishListService) UpdateWishList(
	wishList *model.WishList, updateMask []string,
) (*model.WishList, *exception.ApiError) {
	// TODO: Implement function.
	return nil, exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}

func (s *WishListService) DeleteWishList(
	id int,
) *exception.ApiError {
	// TODO: Implement function.
	return exception.NewApiError(
		http.StatusNotImplemented, "not implemented yet",
	)
}
