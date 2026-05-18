package service

import (
	"net/http"

	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
	"github.com/coding-shenanigans/alchemist-service/internal/repository"
)

type WishListService struct {
	userRepository     *repository.UserRepository
	wishListRepository *repository.WishListRepository
}

func NewWishListService(
	userRepository *repository.UserRepository,
	wishListRepository *repository.WishListRepository,
) *WishListService {
	return &WishListService{
		userRepository:     userRepository,
		wishListRepository: wishListRepository,
	}
}

func (s *WishListService) CreateWishList(
	authenticatedUserId int, username string, wishList *model.WishList,
) (*model.WishList, *exception.ApiError) {
	user, apiErr := s.userRepository.GetUserByUsername(username)
	if apiErr != nil {
		return nil, apiErr
	}

	if user.Id != authenticatedUserId {
		return nil, exception.NewApiError(
			http.StatusForbidden,
			"you do not have permission to modify this user's data",
		)
	}

	wishList.UserId = authenticatedUserId

	newWishList, apiErr := s.wishListRepository.CreateWishList(wishList)
	if apiErr != nil {
		return nil, apiErr
	}

	return newWishList, nil
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
