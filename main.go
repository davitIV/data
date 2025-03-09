package main

import (
	"fmt"
	"io/ioutil"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber HTTP სერვერი
	app := fiber.New()

	// "/" ბმული
	app.Get("/", func(c *fiber.Ctx) error {
		// HTML ფაილის წაკითხვა
		data, err := ioutil.ReadFile("index.html")
		if err != nil {
			return c.Status(500).SendString("Error reading HTML file")
		}

		// HTML ფაილის გადაცემაც
		return c.SendString(string(data))
	})

	fmt.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}
