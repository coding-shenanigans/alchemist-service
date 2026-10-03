package service

import (
	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
	"github.com/coding-shenanigans/alchemist-service/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUserProfile(
	username string,
) (*model.User, *exception.ApiError) {
	user, apiErr := s.userRepository.GetUserByUsername(username)
	if apiErr != nil {
		return nil, apiErr
	}

	return user, nil
}

func (s *UserService) ListUsers(
	q string, pageSize int, lastId int,
) ([]*model.User, bool, *exception.ApiError) {
	// `pageSize+1` is used to check if there is a next page.
	users, apiErr := s.userRepository.ListUsers(q, pageSize+1, lastId)
	if apiErr != nil {
		return nil, false, apiErr
	}

	hasNextPage := false
	if len(users) > pageSize {
		hasNextPage = true
		users = users[:pageSize]
	}

	return users, hasNextPage, nil
}
