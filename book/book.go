package book

import (
	"GoFiberRestApi/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"net/http"
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
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString("Invalid body data")
	}
	if book.Title == "" {
		return c.Status(503).SendString("Invalid body data")
	}

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(http.StatusNotFound).SendString("No book found with id " + id)
	}
	db.Delete(&book)
	return c.Status(http.StatusNoContent).SendString("Book deleted successfully")
}
