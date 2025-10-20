package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 1,
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Second * 1,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 5)
		return c.JSON("Hello, World!")
	})

	if err := app.Listen("localhost:8000"); err != nil {
		panic(err)
	}

}
