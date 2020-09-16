package book

import (
	"GoFiberRestApi/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, bookId)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New Book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete Book")
}
