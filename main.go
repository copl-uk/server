package main

import (
	"log"

	"github.com/copl-uk/server/handlers"
	"github.com/copl-uk/server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	utils.LoadEnv()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	handlers.Setup(app)

	log.Fatal(app.Listen(":" + utils.GetEnv("PORT")))
}
