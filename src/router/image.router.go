package router

import (
	"github.com/Leonardo-Antonio/api-save.image/src/handler"
	"github.com/gofiber/fiber/v2"
)

func Image(app *fiber.App) {
	handler := new(handler.ImageForm)
	group := app.Group("/api/v1/images")
	group.Post("/", handler.SaveImage)
	group.Static("/", "./public/images")
}
