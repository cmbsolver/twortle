package routes

import (
	"twortle/api"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/search", api.SearchWordsHandler)
	apiGroup.Post("/draw", api.DrawPatternHandler)
	apiGroup.Post("/play", api.GetGameWordHandler)
	apiGroup.Post("/check", api.CheckGuessHandler)
	apiGroup.Get("/lengths", api.GetWordLengthsHandler)
	apiGroup.Post("/wordsearch/generate", api.GenerateWordSearchHandler)
	apiGroup.Post("/wordsearch/save", api.SaveWordSearchHandler)
	apiGroup.Get("/wordsearch/list", api.ListWordSearchesHandler)
	apiGroup.Get("/wordsearch/:id", api.GetWordSearchHandler)
	apiGroup.Delete("/wordsearch/:id", api.DeleteWordSearchHandler)
}
