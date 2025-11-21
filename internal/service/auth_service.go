package service

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/coding-shenanigans/alchemist-service/internal/auth"
	"github.com/coding-shenanigans/alchemist-service/internal/dto"
	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/repository"
)

type AuthService struct {
	userRepository    *repository.UserRepository
	sessionRepository *repository.SessionRepository
}

func NewAuthService(
	userRepository *repository.UserRepository,
	sessionRepository *repository.SessionRepository,
) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (s *AuthService) Signup(
	email string, username string, password string,
) (*dto.UserSession, *exception.ApiError) {
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

	accessToken, err := auth.GenerateAccessToken(user.Id)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to generate access token",
		)
	}

	sessionCookie, err := auth.CreateSessionCookie(user.Id)
	if err != nil {
		// TODO: log error
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to create session cookie",
		)
	}

	_, apiErr = s.sessionRepository.CreateSession(user.Id, sessionCookie.Value)
	if apiErr != nil {
		return nil, apiErr
	}

	return &dto.UserSession{
		Email:         user.Email,
		Username:      user.Username,
		AccessToken:   accessToken,
		SessionCookie: sessionCookie,
	}, nil
}
