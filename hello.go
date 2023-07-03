package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"rsc.io/quote"
)

type response struct {
	quote string
}

func main() {
	app := fiber.New()

	app.Get("/api/quote", func(c *fiber.Ctx) error {
		data := response{
			quote: quote.Glass(),
		}

		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
