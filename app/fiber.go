package app

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
)

func CreateFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "ARC",
		BodyLimit:         20 * 1024 * 1024,
		ErrorHandler:      errorHandler,
		StreamRequestBody: true,
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
	})
	return app
}

func errorHandler(ctx fiber.Ctx, err error) error {
	return err
}
