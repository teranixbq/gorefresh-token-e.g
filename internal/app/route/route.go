package route

import (
	"refresh/internal/user/handler"
	"refresh/internal/app/middleware"
	"refresh/pkg/auth"
	"refresh/internal/user/repository"
	"refresh/internal/user/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteInit(f *fiber.App, db *gorm.DB) {
	a := auth.Token{}
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, a)
	userHandler := handler.NewUserhandler(userService,a)

	f.Post("/register", userHandler.Register)
	f.Post("/login", userHandler.Login)
	f.Post("/refresh", userHandler.RefreshToken)
	f.Get("/profile", middleware.JWTMiddleware(), userHandler.GetProfile)
}
