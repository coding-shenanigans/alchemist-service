package service

import (
	"net/http"

	"github.com/coding-shenanigans/alchemist-service/internal/constant"
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

	wishList.UserId = user.Id

	newWishList, apiErr := s.wishListRepository.CreateWishList(wishList)
	if apiErr != nil {
		return nil, apiErr
	}

	return newWishList, nil
}

func (s *WishListService) GetWishList(
	authenticatedUserId int, username string, wishListId int,
) (*model.WishList, *exception.ApiError) {
	user, apiErr := s.userRepository.GetUserByUsername(username)
	if apiErr != nil {
		return nil, apiErr
	}

	wishList, apiErr := s.wishListRepository.GetWishListById(
		user.Id, wishListId,
	)
	if apiErr != nil {
		return nil, apiErr
	}

	// TODO: Check actual friendship when that feature is added.
	isFriend := false

	if !s.hasAccess(authenticatedUserId, wishList, isFriend) {
		return nil, exception.NewApiError(
			http.StatusNotFound, "the wish list was not found",
		)
	}

	return wishList, nil
}

func (s *WishListService) ListWishLists(
	authenticatedUserId int, username string,
) ([]*model.WishList, *exception.ApiError) {
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

	wishLists, apiErr := s.wishListRepository.ListWishLists(user.Id)
	if apiErr != nil {
		return nil, apiErr
	}

	return wishLists, nil
}

func (s *WishListService) UpdateWishList(
	authenticatedUserId int,
	username string,
	wishListId int,
	name *string,
	visibility *string,
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

	wishList, apiErr := s.wishListRepository.GetWishListById(
		user.Id, wishListId,
	)
	if apiErr != nil {
		return nil, apiErr
	}

	if name != nil {
		wishList.Name = *name
	}

	if visibility != nil {
		wishList.Visibility = *visibility
	}

	wishList, apiErr = s.wishListRepository.UpdateWishList(
		user.Id, wishList,
	)
	if apiErr != nil {
		return nil, apiErr
	}

	return wishList, nil
}

func (s *WishListService) DeleteWishList(
	authenticatedUserId int, username string, wishListId int,
) *exception.ApiError {
	user, apiErr := s.userRepository.GetUserByUsername(username)
	if apiErr != nil {
		return apiErr
	}

	if user.Id != authenticatedUserId {
		return exception.NewApiError(
			http.StatusForbidden,
			"you do not have permission to modify this user's data",
		)
	}

	apiErr = s.wishListRepository.DeleteWishListById(user.Id, wishListId)
	if apiErr != nil {
		return apiErr
	}

	return nil
}

// Checks if a user has access to a wish list.
func (s *WishListService) hasAccess(
	authenticatedUserId int, wishList *model.WishList, isFriend bool,
) bool {
	if wishList.Visibility == constant.WishListVisibilityPublic {
		return true
	}

	if wishList.Visibility == constant.WishListVisibilityFriendsOnly && isFriend {
		return true
	}

	return wishList.UserId == authenticatedUserId
}
