package main

import (
	"github.com/ahmedkhalaf1996/Vuityfy-Go-App/routes"

	"github.com/ahmedkhalaf1996/Vuityfy-Go-App/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	models.ConnectDatabase() // New

	app := fiber.New(fiber.Config{
		AppName:      "Fiber with Planetscale",
		ServerHeader: "Fiber",
	})
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080, http://localhost:8080/",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	// setup routes
	routes.Setup(app)

	app.Listen(":5000")
}
