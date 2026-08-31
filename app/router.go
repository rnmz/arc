package app

import "github.com/gofiber/fiber/v3"

func router(app *fiber.App) {
	api := app.Group("/api", func() {})

	api.Get("/file/get", func(c *fiber.Ctx) {})
	api.Post("/file/upload", func(c *fiber.Ctx) {})
	api.Delete("/file/delete", func(c *fiber.Ctx) {})

	api.Post("/folder/create", func(c *fiber.Ctx) {})
	api.Post("/folder/update", func(c *fiber.Ctx) {})
	api.Post("/folder/delete", func(c *fiber.Ctx) {})
	api.Get("/folder/list", func(c *fiber.Ctx) {})
	api.Get("/folder/get", func(c *fiber.Ctx) {})

	auth := app.Group("/auth")
	auth.Post("/login", func(c *fiber.Ctx) {})
	auth.Post("/logout", func(c *fiber.Ctx) {})
	auth.Post("/register", func(c *fiber.Ctx) {})
	auth.Post("/recover", func(c *fiber.Ctx) {})

	auth.Post("/key/refresh", func(c *fiber.Ctx) {})
	auth.Post("/key/generate", func(c *fiber.Ctx) {})
	auth.Post("/key/revoke", func(c *fiber.Ctx) {})
}
