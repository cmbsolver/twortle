package api

import (
	"strings"
	"twortle/pkg/logic"

	"github.com/gofiber/fiber/v2"
)

type SearchRequest struct {
	Text     string   `json:"text"`
	Contains []string `json:"contains"`
	Exclude  []string `json:"exclude"`
}

func SearchWordsHandler(c *fiber.Ctx) error {
	var req SearchRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	results, err := logic.SearchWordleWords(strings.Split(req.Text, ""), req.Contains, req.Exclude)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(results)
}
