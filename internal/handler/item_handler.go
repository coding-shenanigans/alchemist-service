package handler

import (
	"net/http"
	"strconv"

	"github.com/coding-shenanigans/alchemist-service/internal/constant"
	"github.com/coding-shenanigans/alchemist-service/internal/dto"
	"github.com/coding-shenanigans/alchemist-service/internal/service"
	"github.com/gin-gonic/gin"
)

type itemHandler struct {
	itemService *service.ItemService
}

func newItemHandler(
	itemService *service.ItemService,
) *itemHandler {
	return &itemHandler{
		itemService: itemService,
	}
}

func (h *itemHandler) createItem(c *gin.Context) {
	req := new(dto.CreateItemRequest)

	if err := c.ShouldBindJSON(req); err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, err.Error()))
		return
	}

	authenticatedUserId := c.GetInt(constant.AuthenticatedUserId)
	username := c.Param("username")
	wishListIdParam := c.Param("wishListId")
	wishListId, err := strconv.Atoi(wishListIdParam)
	if err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, "invalid wish list id"))
		return
	}

	req.Item.WishListId = wishListId

	item, apiErr := h.itemService.CreateItem(
		authenticatedUserId, username, req.Item,
	)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	res := &dto.CreateItemResponse{Item: item}
	c.JSON(http.StatusCreated, res)
}

func (h *itemHandler) getItem(c *gin.Context) {
	authenticatedUserId := c.GetInt(constant.AuthenticatedUserId)
	username := c.Param("username")

	wishListId, err := strconv.Atoi(c.Param("wishListId"))
	if err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, "invalid wish list id"))
		return
	}

	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, "invalid item id"))
		return
	}

	item, apiErr := h.itemService.GetItem(
		authenticatedUserId, username, wishListId, itemId,
	)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	res := &dto.GetItemResponse{Item: item}
	c.JSON(http.StatusOK, res)
}

func (h *itemHandler) listItems(c *gin.Context) {
	authenticatedUserId := c.GetInt(constant.AuthenticatedUserId)
	username := c.Param("username")

	wishListId, err := strconv.Atoi(c.Param("wishListId"))
	if err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, "invalid wish list id"))
		return
	}

	items, apiErr := h.itemService.ListItems(
		authenticatedUserId, username, wishListId,
	)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	res := &dto.ListItemsResponse{Items: items}
	c.JSON(http.StatusOK, res)
}

func (h *itemHandler) deleteItem(c *gin.Context) {
	authenticatedUserId := c.GetInt(constant.AuthenticatedUserId)
	username := c.Param("username")

	wishListId, err := strconv.Atoi(c.Param("wishListId"))
	if err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, "invalid wish list id"))
		return
	}

	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, "invalid item id"))
		return
	}

	apiErr := h.itemService.DeleteItem(
		authenticatedUserId, username, wishListId, itemId,
	)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
