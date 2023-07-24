package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	dbmodels "xfiber/dbModels"
	"xfiber/storage"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()
	app.Use(cors.New())
	// Create a new PostgreSQL database connection
	client, err := storage.NewPostgreSQLConnection()
	if err != nil {
		// Log a fatal error if the connection cannot be established
		log.Fatalf("failed to create sql client: %s", err)
	}

	// Defer the closing of the database connection to ensure that it is always closed
	defer func(client *sqlx.DB) {
		err := client.Close()
		if err != nil {
			// Log a fatal error if the connection cannot be closed
			log.Fatalf("failed to close postgres client: %s", err)
		}
	}(client)

	err = client.Ping()
	if err != nil {
		log.Fatalf("No ping from db")
	}

	// Create new GET route on path "/hello"
	app.Get("/tenants", func(c *fiber.Ctx) error {
		tenants, err := dbmodels.Tenants().All(context.Background(), client)
		if err != nil {
			c.Status(404)
		}
		return c.Status(200).JSON(tenants)
	})

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
