package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// ფერის ტექსტი კონსოლში
	red := color.New(color.FgRed).Add(color.Underline)
	blue := color.New(color.FgBlue).Add(color.Bold)

	// ლექსი სიყვარულზე
	poem := `
	ერთ წამში ფიქრებში დავხრჩიე,
	სიყვარულში გალაღებული გული,
	რომ წამოიძახო ჩემი სახელი,
	ეს არის ჩემი ცხოვრების მთავარი მიზანი. 💖
	`

	// კონსოლში ლექსის ფერით გაშვება
	red.Println("💖 " + poem)

	// Fiber HTTP სერვერი
	app := fiber.New()

	// "/" ბმული
	app.Get("/", func(c *fiber.Ctx) error {
		// HTML გვერდზე ლექსი, გამოყენებული HTML ფორმატირება
		return c.SendString(`
		<!DOCTYPE html>
		<html lang="ka">
		<head>
			<meta charset="UTF-8">
			<title>სიყვარულის ლექსი</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f8f8ff;
					text-align: center;
					padding: 50px;
				}
				h1 {
					color: #ff1493;
					font-size: 3em;
					text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
				}
				.poem {
					color: #8b0000;
					font-size: 1.5em;
					font-weight: bold;
					margin-top: 20px;
					line-height: 1.8;
				}
				.poem span {
					color: #4682b4;
					font-style: italic;
				}
			</style>
		</head>
		<body>
			<h1>💖 ერთ წამში ფიქრებში დავხრჩიე 💖</h1>
			<p class="poem">
				💖 <span>ერთ წამში ფიქრებში დავხრჩიე,</span> <br>
				სიყვარულში გალაღებული გული,<br>
				რომ წამოიძახო ჩემი სახელი,<br>
				ეს არის ჩემი ცხოვრების მთავარი მიზანი. 💕<br>
			</p>
		</body>
		</html>
		`)
	})

	fmt.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}
