package main

import (
	"main/bootstrapper"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	err := bootstrapper.New().RunAPI()

	if err != nil {
		return
	}

	app := fiber.New()

	app.Get("/", get)


	app.Listen(":3000")

}

func get(c *fiber.Ctx) error {
	return c.JSON("Hello, World!")
}
