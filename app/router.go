package app

import "github.com/gofiber/fiber/v3"

func router(app *fiber.App) {
	file := app.Group("/files")
	file.Get("/get", func(c *fiber.Ctx) {})
	file.Post("/upload", func(c *fiber.Ctx) {})
	file.Post("/move", func(c *fiber.Ctx) {})
	file.Delete("/delete", func(c *fiber.Ctx) {})

	folder := app.Group("/folder")
	folder.Post("/create", func(c *fiber.Ctx) {})
	folder.Post("/update", func(c *fiber.Ctx) {})
	folder.Post("/delete", func(c *fiber.Ctx) {})
	folder.Get("/list", func(c *fiber.Ctx) {})
	folder.Get("/get", func(c *fiber.Ctx) {})

	auth := app.Group("/auth")
	auth.Post("/login", func(c *fiber.Ctx) {})
	auth.Post("/logout", func(c *fiber.Ctx) {})
	auth.Post("/register", func(c *fiber.Ctx) {})
	auth.Post("/recover", func(c *fiber.Ctx) {})

	auth.Post("/key/refresh", func(c *fiber.Ctx) {})
	auth.Post("/key/generate", func(c *fiber.Ctx) {})
	auth.Post("/key/revoke", func(c *fiber.Ctx) {})

	account := app.Group("/account")
	account.Post("/profile", func(c *fiber.Ctx) {})
	account.Post("/change/password", func(c *fiber.Ctx) {})
	account.Post("/change/email", func(c *fiber.Ctx) {})
	account.Post("/validate/email", func(c *fiber.Ctx) {})

	admin := app.Group("/admin")
	admin.Get("/user/get", func(c *fiber.Ctx) {})
	admin.Get("/users/get", func(c *fiber.Ctx) {})
	admin.Post("/change/role", func(c *fiber.Ctx) {})
	admin.Post("/change/plan", func(c *fiber.Ctx) {})
	admin.Get("/system", func(c *fiber.Ctx) {})

	app.Post("/smtp/send", func(c *fiber.Ctx) {})
}
