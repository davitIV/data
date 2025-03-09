package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// ფერის ტექსტი კონსოლში
	red := color.New(color.FgRed)
	red.Println("გამარჯობა სიდო 🚀")

	// Fiber HTTP სერვერი
	app := fiber.New()

	// "/" ბმული
	app.Get("/", func(c *fiber.Ctx) error {
		// HTML გვერდი, სადაც ტექსტი წითლად იქნება
		return c.SendString("<h1 style='color:red;'>გამარჯობა სიდო 🚀</h1>")
	})

	fmt.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}
