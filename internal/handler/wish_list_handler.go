package handler

import (
	"net/http"

	"github.com/coding-shenanigans/alchemist-service/internal/dto"
	"github.com/coding-shenanigans/alchemist-service/internal/service"
	"github.com/gin-gonic/gin"
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
	// TODO: Implement function.
	status := http.StatusNotImplemented
	c.JSON(status, dto.NewErrorResponse(status, "not implemented yet"))
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
