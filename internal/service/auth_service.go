package service

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
	"github.com/coding-shenanigans/alchemist-service/internal/repository"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Signup(
	email string, username string, password string,
) (*model.User, *exception.ApiError) {
	apiErr := s.userRepository.EmailExists(email)
	if apiErr != nil {
		return nil, apiErr
	}

	apiErr = s.userRepository.UsernameExists(username)
	if apiErr != nil {
		return nil, apiErr
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to hash the password",
		)
	}

	user, apiErr := s.userRepository.CreateUser(
		email, username, string(hashedPasswordBytes),
	)
	if apiErr != nil {
		return nil, apiErr
	}

	return user, nil
}
