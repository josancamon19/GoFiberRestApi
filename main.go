package main

import "github.com/gofiber/fiber"

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
}

func main() {
	app := fiber.New()
	app.Listen(":3000")
}
