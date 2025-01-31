package main

import (
	"fmt"
	"main/bootstrapper"

	"github.com/gofiber/fiber/v2"
)

func main() {
	bootstrapper.New().RunAPI()

	app := fiber.New()

	app.Get("/", GetNews)

	app.Listen(":8080")

}

func GetNews(c *fiber.Ctx) error {
	fmt.Println("Hello, World!")
	return c.JSON("Hello, World!")
}
