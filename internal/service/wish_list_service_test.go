package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coding-shenanigans/alchemist-service/internal/constant"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
)

func TestHasAccess(t *testing.T) {
	s := &WishListService{}

	tests := []struct {
		name                string
		authenticatedUserId int
		wishList            *model.WishList
		isFriend            bool
		expected            bool
	}{
		{
			name:                "unauthenticated_can_access_public_wish_list",
			authenticatedUserId: 0,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityPublic,
			},
			isFriend: false,
			expected: true,
		},
		{
			name:                "friend_user_can_access_public_wish_list",
			authenticatedUserId: 2,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityPublic,
			},
			isFriend: true,
			expected: true,
		},
		{
			name:                "owner_can_access_public_wish_list",
			authenticatedUserId: 1,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityPublic,
			},
			isFriend: false,
			expected: true,
		},
		{
			name:                "unauthenticated_user_cannot_access_friends_only_wish_list",
			authenticatedUserId: 0,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityFriendsOnly,
			},
			isFriend: false,
			expected: false,
		},
		{
			name:                "non_friend_user_cannot_access_friends_only_wish_list",
			authenticatedUserId: 3,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityFriendsOnly,
			},
			isFriend: false,
			expected: false,
		},
		{
			name:                "friend_user_can_access_friends_only_wish_list",
			authenticatedUserId: 2,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityFriendsOnly,
			},
			isFriend: true,
			expected: true,
		},
		{
			name:                "owner_can_access_friends_only_wish_list",
			authenticatedUserId: 1,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityFriendsOnly,
			},
			isFriend: false,
			expected: true,
		},
		{
			name:                "unauthenticated_user_cannot_access_private_wish_list",
			authenticatedUserId: 0,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityPrivate,
			},
			isFriend: false,
			expected: false,
		},
		{
			name:                "friend_user_cannot_access_private_wish_list",
			authenticatedUserId: 2,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityPrivate,
			},
			isFriend: true,
			expected: false,
		},
		{
			name:                "owner_can_access_private_wish_list",
			authenticatedUserId: 1,
			wishList: &model.WishList{
				UserId:     1,
				Visibility: constant.WishListVisibilityPrivate,
			},
			isFriend: false,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := s.hasAccess(tt.authenticatedUserId, tt.wishList, tt.isFriend)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFilterAccessibleWishLists(t *testing.T) {
	s := &WishListService{}

	tests := []struct {
		name                string
		authenticatedUserId int
		wishLists           []*model.WishList
		isFriend            bool
		expected            []*model.WishList
	}{
		{
			name:                "unauthenticated_can_access_public_wish_lists",
			authenticatedUserId: 0,
			wishLists: []*model.WishList{
				{UserId: 1, Visibility: constant.WishListVisibilityPublic},
				{UserId: 1, Visibility: constant.WishListVisibilityFriendsOnly},
				{UserId: 1, Visibility: constant.WishListVisibilityPrivate},
			},
			isFriend: false,
			expected: []*model.WishList{
				{UserId: 1, Visibility: constant.WishListVisibilityPublic},
			},
		},
		{
			name:                "friend_user_can_access_friends_only_wish_lists",
			authenticatedUserId: 2,
			wishLists: []*model.WishList{
				{UserId: 1, Visibility: constant.WishListVisibilityPublic},
				{UserId: 1, Visibility: constant.WishListVisibilityFriendsOnly},
				{UserId: 1, Visibility: constant.WishListVisibilityPrivate},
			},
			isFriend: true,
			expected: []*model.WishList{
				{UserId: 1, Visibility: constant.WishListVisibilityPublic},
				{UserId: 1, Visibility: constant.WishListVisibilityFriendsOnly},
			},
		},
		{
			name:                "owner_can_access_private_wish_lists",
			authenticatedUserId: 1,
			wishLists: []*model.WishList{
				{UserId: 1, Visibility: constant.WishListVisibilityPublic},
				{UserId: 1, Visibility: constant.WishListVisibilityFriendsOnly},
				{UserId: 1, Visibility: constant.WishListVisibilityPrivate},
			},
			isFriend: true,
			expected: []*model.WishList{
				{UserId: 1, Visibility: constant.WishListVisibilityPublic},
				{UserId: 1, Visibility: constant.WishListVisibilityFriendsOnly},
				{UserId: 1, Visibility: constant.WishListVisibilityPrivate},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := s.filterAccessibleWishLists(
				tt.authenticatedUserId, tt.wishLists, tt.isFriend,
			)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
