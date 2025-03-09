package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// ფერის ტექსტი კონსოლში
	red := color.New(color.FgRed).Add(color.Underline)
	red.Println("გამარჯობა სიდო 🚀")

	// Fiber HTTP სერვერი
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("გამარჯობა სიდო 🚀")
	})

	fmt.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}
