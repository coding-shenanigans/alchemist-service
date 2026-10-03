package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/domain"
	"github.com/coding-shenanigans/alchemist-service/internal/dto"
	"github.com/coding-shenanigans/alchemist-service/internal/service"
)

type userHandler struct {
	userService *service.UserService
}

func newUserHandler(userService *service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) getUserProfile(c *gin.Context) {
	username := c.Param("username")

	user, apiErr := h.userService.GetUserProfile(username)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	res := &dto.GetUserProfileResponse{Username: user.Username}
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) listUsers(c *gin.Context) {
	q := c.Query("q")

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, "invalid page size"))
		return
	}

	pageToken := c.Query("pageToken")
	pageCursor := &domain.UserPageCursor{Q: q, PageSize: pageSize, LastId: 0}
	if pageToken != "" {
		pageCursor, err = decodePageToken(pageToken)
		if err != nil {
			status := http.StatusBadRequest
			c.JSON(status, dto.NewErrorResponse(status, "invalid page token"))
			return
		}

		if pageCursor.Q != q || pageCursor.PageSize != pageSize {
			status := http.StatusBadRequest
			errMessage := "page token doesn't match the query params"
			c.JSON(status, dto.NewErrorResponse(status, errMessage))
			return
		}
	}

	users, hasNextPage, apiErr := h.userService.ListUsers(
		pageCursor.Q, pageCursor.PageSize, pageCursor.LastId,
	)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	userResults := make([]*dto.UserResult, len(users))
	for i, user := range users {
		userResults[i] = &dto.UserResult{Username: user.Username}
		pageCursor.LastId = user.Id
	}

	res := &dto.ListUsersResponse{Users: userResults}
	if hasNextPage {
		nextPageToken, err := encodePageToken(pageCursor)
		if err != nil {
			status := http.StatusInternalServerError
			errMessage := "failed to create next page token"
			c.JSON(status, dto.NewErrorResponse(status, errMessage))
			return
		}

		res.NextPageToken = nextPageToken
	}

	c.JSON(http.StatusOK, res)
}

func encodePageToken(pageCursor *domain.UserPageCursor) (string, error) {
	data, err := json.Marshal(pageCursor)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(data), nil
}

func decodePageToken(pageToken string) (*domain.UserPageCursor, error) {
	data, err := base64.RawURLEncoding.DecodeString(pageToken)
	if err != nil {
		return nil, err
	}

	param := new(domain.UserPageCursor)
	if err := json.Unmarshal(data, param); err != nil {
		return nil, err
	}

	return param, nil
}
