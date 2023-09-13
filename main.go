package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/pusher/pusher-http-go/v5"
)

// pusher id= lavat70879@horsgit.com

func main() {
	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1665310",
		Key:     "6d32bd78ee0bd5e4ebac",
		Secret:  "eb40edcc1767b045aa18",
		Cluster: "ap2",
		Secure:  true,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Brother 👋!")
	})

	app.Post("/api/messages", func(c *fiber.Ctx) error {

		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chatbox", "message", data)
		return c.SendString("Message sent!!")
	})

	app.Listen(":3000")
}
