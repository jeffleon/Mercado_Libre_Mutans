package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffleon/mutansApi/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/mutant", handler.PostMutant)
	v1.Get("/stats/:id", handler.GetStats)
	v1.Get("/stats", handler.GetSAlltats)
}