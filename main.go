package main

import (
	"log"

	"github.com/fadhlimulyana20/sister-jwt/database"
	"github.com/fadhlimulyana20/sister-jwt/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]interface{}{
			"hello": "world",
		})
	})

	api := new(router.Api)
	api.Init(app)

	log.Fatal(app.Listen(":5500"))
}
