package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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
	viper.SetConfigFile("ENV")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	port := fmt.Sprint(viper.Get("PORT"))

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

	app.Listen("0.0.0.0:" + port)
}
