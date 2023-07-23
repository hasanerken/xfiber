package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Create new GET route on path "/hello"
	app.Get("/check", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))

	//   DATABASE_URL=postgres://xfiber:nay4LTqtq4epj9Q@xfiber-db.flycast:5432/xfiber?sslmode=disable
}
