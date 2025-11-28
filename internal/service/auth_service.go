package service

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/coding-shenanigans/alchemist-service/internal/auth"
	"github.com/coding-shenanigans/alchemist-service/internal/dto"
	"github.com/coding-shenanigans/alchemist-service/internal/exception"
	"github.com/coding-shenanigans/alchemist-service/internal/model"
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

	return s.createUserSession(user)
}

func (s *AuthService) Signin(
	email string, password string,
) (*dto.UserSession, *exception.ApiError) {
	user, apiErr := s.userRepository.GetUserByEmail(email)
	if apiErr != nil {
		if apiErr.Status() == http.StatusNotFound {
			return nil, exception.NewApiError(
				http.StatusUnauthorized, "invalid email or password",
			)
		} else {
			return nil, apiErr
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, exception.NewApiError(
			http.StatusUnauthorized, "invalid email or password",
		)
	}

	return s.createUserSession(user)
}

func (s *AuthService) Refresh(
	refreshToken string,
) (*dto.UserSession, *exception.ApiError) {
	_, err := auth.ValidateToken(refreshToken)
	if err != nil {
		return nil, exception.NewApiError(http.StatusUnauthorized, err.Error())
	}

	session, apiErr := s.sessionRepository.GetSessionByRefreshToken(refreshToken)
	if apiErr != nil {
		return nil, apiErr
	}

	user, apiErr := s.userRepository.GetUserById(session.UserId)
	if apiErr != nil {
		return nil, apiErr
	}

	return s.refreshUserSession(user, session)
}

func (s *AuthService) Signout(refreshToken string) *exception.ApiError {
	apiErr := s.sessionRepository.DeleteSessionByRefreshToken(refreshToken)
	if apiErr != nil {
		return apiErr
	}

	return nil
}

// Creates a user session.
func (s *AuthService) createUserSession(
	user *model.User,
) (*dto.UserSession, *exception.ApiError) {
	accessToken, err := auth.GenerateAccessToken(user.Id)
	if err != nil {
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to generate access token",
		)
	}

	refreshToken, err := auth.GenerateRefreshToken(user.Id)
	if err != nil {
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to generate refresh token",
		)
	}

	_, apiErr := s.sessionRepository.CreateSession(user.Id, refreshToken)
	if apiErr != nil {
		return nil, apiErr
	}

	return &dto.UserSession{
		Email:        user.Email,
		Username:     user.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Refreshes a user session.
func (s *AuthService) refreshUserSession(
	user *model.User, session *model.Session,
) (*dto.UserSession, *exception.ApiError) {
	accessToken, err := auth.GenerateAccessToken(user.Id)
	if err != nil {
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to generate access token",
		)
	}

	refreshToken, err := auth.GenerateRefreshToken(user.Id)
	if err != nil {
		return nil, exception.NewApiError(
			http.StatusInternalServerError, "failed to generate refresh token",
		)
	}

	apiErr := s.sessionRepository.RefreshSession(session.Id, refreshToken)
	if apiErr != nil {
		return nil, apiErr
	}

	return &dto.UserSession{
		Email:        user.Email,
		Username:     user.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
