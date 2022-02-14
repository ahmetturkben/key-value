package router

import (
	"record/key/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func MemoryRouter(app *fiber.App, handler *http.MemoryHandler) {
	api := app.Group("/api")
	inmemory := api.Group("/memory")
	inmemory.Get("/:key", handler.GetMemoryByKey)
	inmemory.Get("/all/flush", handler.Flush)
	inmemory.Post("/", handler.CreateKey)
}
