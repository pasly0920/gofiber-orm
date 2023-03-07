package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-orm/app/http/handler"
)

func UserRouter(app fiber.Router, handler *handler.UserHandler) {
	app.Get("/user/:id", handler.GetUser)
}
