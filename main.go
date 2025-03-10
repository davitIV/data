package main

import (
	"context"
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
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// context.Context-ი საჭიროა PostgreSQL კავშირის დასამყარებლად
	ctx := context.Background()

	var err error
	db, err = pgxpool.New(ctx, databaseUrl)
	if err != nil {
		log.Fatal("DB Connection Error: ", err)
	}
	defer db.Close()

	app := fiber.New()

	// "/" ბმული
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// კავშირის წარმატებით დამყარების შეტყობინება
	fmt.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}
