package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"rsc.io/quote"
)

type SomeQuote struct {
	Quote string
}

type SomeStruct struct {
	Name string
	Age  uint8
}

func main() {
	app := fiber.New()

	app.Get("/api/quote", func(c *fiber.Ctx) error {
		quote := quote.Glass()
		data := SomeQuote{
			Quote: quote,
		}

		return c.JSON(data)
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		// Create data struct:
		data := SomeStruct{
			Name: "Grame",
			Age:  20,
		}

		return c.JSON(data)
	})

	log.Println(apo.ListenAndServe(":"+port)
}
