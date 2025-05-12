package main

import (
	"log"
	"os"

	"github.com/copl-uk/server/handlers"
	"github.com/copl-uk/server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	utils.SetupEnv()
}

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	handlers.Setup(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
