package main

/*
Carrega config.
Conecta no DB.
Sobe Fiber.
Registra handlers da API.
*/

import (
	"github.com/gofiber/fiber/v2"
	"github.com/psousaj/nfcpdv/config"
	"github.com/psousaj/nfcpdv/internal/infra/db"
)

func main() {
	app := fiber.New()
	config := config.LoadConfig()
	_ = db.ConnectPostgres(config.DBUrl)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Listen(":3000")
}
