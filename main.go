package main

import "github.com/gofiber/fiber"

//import "github.com/josancamon19/GoFiberRestApi/book"

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
