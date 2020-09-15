package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Single Book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New Book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete Book")
}
