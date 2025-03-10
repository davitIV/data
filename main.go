package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func main() {
	// PostgreSQL კავშირის შექმნა
	databaseUrl := os.Getenv("DATABASE_URL") // Railway-სგან მიღებული URL
	var err error
	db, err = pgxpool.New(nil, databaseUrl)
	if err != nil {
		log.Fatal("DB Connection Error: ", err)
	}
	defer db.Close()

	app := fiber.New()

	// "/" ბმული
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}
