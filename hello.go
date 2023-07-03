package main

import (
	"github.com/gofiber/fiber"
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
}
