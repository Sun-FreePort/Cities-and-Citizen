package router

import (
	"github.com/Sun-FreePort/Cities-and-Citizen/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	H *handler.Handler
}

func (r Router) RegisterF2E(app *fiber.App) {
	app.Static("/", "./public")

	app.Get("/swagger/*", swagger.HandlerDefault) // default
}

func (r Router) RegisterB2E(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Ops! nothing happen...")
	})

	app.Get("/square/info", r.H.SquareInfo)
}
