package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/database"
	"github.com/coding-shenanigans/alchemist-service/internal/repository"
	"github.com/coding-shenanigans/alchemist-service/internal/service"
)

// Registers all the endpoints handled by the service.
func RegisterEndpoints(public *gin.RouterGroup, protected *gin.RouterGroup) {
	// connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// create repositories
	userRepository := repository.NewUserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	wishListRepository := repository.NewWishListRepository(db)

	// create services
	authService := service.NewAuthService(userRepository, sessionRepository)
	wishListService := service.NewWishListService(userRepository, wishListRepository)

	// create handlers
	opsHandler := newOpsHandler()
	authHandler := newAuthHandler(authService)
	wishListHandler := newWishListHandler(wishListService)

	// register endpoints
	public.GET("/health", opsHandler.health)

	public.POST("/auth/signup", authHandler.signup)
	public.POST("/auth/signin", authHandler.signin)
	public.POST("/auth/refresh", authHandler.refresh)
	protected.POST("/auth/signout", authHandler.signout)

	protected.POST("/users/:username/wish-lists", wishListHandler.createWishList)
	protected.GET("/users/:username/wish-lists", wishListHandler.listWishLists)
	protected.GET("/users/:username/wish-lists/:wishListId", wishListHandler.getWishList)
	protected.PATCH("/users/:username/wish-lists/:wishListId", wishListHandler.updateWishList)
	protected.DELETE("/users/:username/wish-lists/:wishListId", wishListHandler.deleteWishList)
}
