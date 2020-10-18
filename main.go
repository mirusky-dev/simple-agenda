package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/mirusky-dev/simple-agenda/database"
	"github.com/mirusky-dev/simple-agenda/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	db, err := database.New(connStr)
	if err != nil {
		log.Fatalf("DATABASE ERROR: %+v", err.Error())
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong!",
		})
	})

	api := app.Group("/api")
	handlers.AppointmentRouter(api, db)

	if port := os.Getenv("PORT"); port != "" {
		app.Listen(":" + port)
	} else {
		app.Listen(":3030")
	}
}
