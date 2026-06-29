package service

import (
	"net/http"

	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
	"github.com/coding-shenanigans/alchemist-service/internal/repository"
)

type ItemService struct {
	userRepository     *repository.UserRepository
	wishListRepository *repository.WishListRepository
	itemRepository     *repository.ItemRepository
}

func NewItemService(
	userRepository *repository.UserRepository,
	wishListRepository *repository.WishListRepository,
	itemRepository *repository.ItemRepository,
) *ItemService {
	return &ItemService{
		userRepository:     userRepository,
		wishListRepository: wishListRepository,
		itemRepository:     itemRepository,
	}
}

func (s *ItemService) CreateItem(
	authenticatedUserId int, username string, item *model.Item,
) (*model.Item, *exception.ApiError) {
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

	_, apiErr = s.wishListRepository.GetWishListById(
		user.Id, item.WishListId,
	)
	if apiErr != nil {
		return nil, apiErr
	}

	newItem, apiErr := s.itemRepository.CreateItem(item)
	if apiErr != nil {
		return nil, apiErr
	}

	return newItem, nil
}

func (s *ItemService) GetItem(
	authenticatedUserId int, username string, wishListId int, itemId int,
) (*model.Item, *exception.ApiError) {
	user, apiErr := s.userRepository.GetUserByUsername(username)
	if apiErr != nil {
		return nil, apiErr
	}

	wishList, apiErr := s.wishListRepository.GetWishListById(user.Id, wishListId)
	if apiErr != nil {
		return nil, apiErr
	}

	// TODO: Friendships have not been implemented yet. When friendships are
	// implemented, we need to properly populate this value.
	isFriend := false

	if !hasAccessToWishList(authenticatedUserId, wishList, isFriend) {
		return nil, exception.NewApiError(
			http.StatusNotFound, "the wish list was not found",
		)
	}

	item, apiErr := s.itemRepository.GetItemById(wishListId, itemId)
	if apiErr != nil {
		return nil, apiErr
	}

	return item, nil
}

func (s *ItemService) ListItems(
	authenticatedUserId int, username string, wishListId int,
) ([]*model.Item, *exception.ApiError) {
	user, apiErr := s.userRepository.GetUserByUsername(username)
	if apiErr != nil {
		return nil, apiErr
	}

	wishList, apiErr := s.wishListRepository.GetWishListById(user.Id, wishListId)
	if apiErr != nil {
		return nil, apiErr
	}

	// TODO: Friendships have not been implemented yet. When friendships are
	// implemented, we need to properly populate this value.
	isFriend := false

	if !hasAccessToWishList(authenticatedUserId, wishList, isFriend) {
		return nil, exception.NewApiError(
			http.StatusNotFound, "the wish list was not found",
		)
	}

	items, apiErr := s.itemRepository.ListItems(wishListId)
	if apiErr != nil {
		return nil, apiErr
	}

	return items, nil
}
