package router

import (
	"github.com/copl-uk/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/ping", handlers.Ping)
}
