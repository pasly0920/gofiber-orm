package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"gofiber-orm/app/http/handler"
	"gofiber-orm/app/http/routes"
	"gofiber-orm/domain/repository"
	"gofiber-orm/domain/service"
	"gofiber-orm/infrastructure/database"
	"gofiber-orm/infrastructure/logger"
	"log"
)

func main() {
	app := fiber.New()

	// Initialize logger
	logger := logger.NewLogger()

	// Set up middleware
	app.Use(logger)
	app.Use(recover2.New())
	app.Use(cors.New())

	//initialize db
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	//Initialize repositories
	userRepository := repository.NewUserRepository(db)

	// Initialize services
	userSvc := service.NewUserService(userRepository)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userSvc, logger)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("")
	})

	api := app.Group("/api")
	routes.UserRouter(api, userHandler)

	app.Listen(":8081")
}
