package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// მთავარი პუნქტი
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(`
			<!DOCTYPE html>
			<html lang="ka">
			<head>
				<meta charset="UTF-8">
				<title>ფერების შეცვლა</title>
				<style>
					body {
						background-color: #f0f0f0;
					}
					h1 {
						color: red; /* აქ შეგიძლიათ შეცვალოთ ფერი */
					}
				</style>
			</head>
			<body>
				<h1>გამარჯობა სიდო 🚀</h1>
			</body>
			</html>
		`)
	})

	fmt.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}
