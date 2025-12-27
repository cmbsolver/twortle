package api

import (
	"twortle/pkg/db/tables"
	"twortle/pkg/logic"

	"github.com/gofiber/fiber/v2"
)

type GameWordRequest struct {
	Length int `json:"length"`
}

type GuessRequest struct {
	Word  tables.Word `json:"word"`
	Guess string      `json:"guess"`
}

type GuessResponse struct {
	Pattern    logic.StringPattern `json:"pattern"`
	IsAllMatch bool                `json:"isAllMatch"`
}

func GetGameWordHandler(c *fiber.Ctx) error {
	var req GameWordRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	result := logic.GetWordForGame(req.Length)

	return c.JSON(result)
}

func CheckGuessHandler(c *fiber.Ctx) error {
	var req GuessRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	result, isAllMatch := logic.CheckWordPattern(req.Guess, req.Word)

	response := GuessResponse{
		Pattern:    result,
		IsAllMatch: isAllMatch,
	}

	return c.JSON(response)
}

func GetWordLengthsHandler(c *fiber.Ctx) error {
	lengths := logic.GetWordLengths()
	return c.JSON(lengths)
}
