package app

import (
	"os"

	"github.com/Leonardo-Antonio/api-save.image/src/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type server struct {
	app *fiber.App
}

func NewServer() *server {
	return &server{app: fiber.New()}
}

func (a *server) Middlewares() {
	a.app.Use(logger.New())
}

func (a *server) Routers() {
	router.Image(a.app)
}

func (a *server) Listeing() {
	a.app.Listen(":" + os.Getenv("PORT"))
}
