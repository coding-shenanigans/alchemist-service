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
