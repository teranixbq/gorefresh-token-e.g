package main

import (
	"refresh/internal/app/config"
	"refresh/internal/app/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	f := fiber.New()

	cfg := config.InitConfig()
	db := config.InitDBPostgres(cfg)
	db.AutoMigrate()

	route.RouteInit(f, db)
	f.Use(cors.New())

	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	f.Listen(":8080")
}
