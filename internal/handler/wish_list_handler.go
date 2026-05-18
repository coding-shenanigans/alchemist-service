package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/constant"
	"github.com/coding-shenanigans/alchemist-service/internal/dto"
	"github.com/coding-shenanigans/alchemist-service/internal/service"
)

type wishListHandler struct {
	wishListService *service.WishListService
}

func newWishListHandler(
	wishListService *service.WishListService,
) *wishListHandler {
	return &wishListHandler{
		wishListService: wishListService,
	}
}

func (h *wishListHandler) createWishList(c *gin.Context) {
	req := new(dto.CreateWishListRequest)

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

	wishList, apiErr := h.wishListService.CreateWishList(
		authenticatedUserId, username, req.WishList,
	)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	res := &dto.CreateWishListResponse{WishList: wishList}
	c.JSON(http.StatusCreated, res)
}

func (h *wishListHandler) getWishList(c *gin.Context) {
	// TODO: Implement function.
	status := http.StatusNotImplemented
	c.JSON(status, dto.NewErrorResponse(status, "not implemented yet"))
}

func (h *wishListHandler) listWishLists(c *gin.Context) {
	// TODO: Implement function.
	status := http.StatusNotImplemented
	c.JSON(status, dto.NewErrorResponse(status, "not implemented yet"))
}

func (h *wishListHandler) updateWishList(c *gin.Context) {
	// TODO: Implement function.
	status := http.StatusNotImplemented
	c.JSON(status, dto.NewErrorResponse(status, "not implemented yet"))
}

func (h *wishListHandler) deleteWishList(c *gin.Context) {
	// TODO: Implement function.
	status := http.StatusNotImplemented
	c.JSON(status, dto.NewErrorResponse(status, "not implemented yet"))
}
