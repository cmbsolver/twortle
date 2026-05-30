package api

import (
	"twortle/pkg/db/tables"
	"twortle/pkg/logic"

	"github.com/gofiber/fiber/v2"
)

// GameWordRequest represents a request for a target word of a specific length.
type GameWordRequest struct {
	Length int `json:"length"`
}

// GuessRequest represents a Wordle guess submission.
type GuessRequest struct {
	Word  tables.Word `json:"word"`
	Guess string      `json:"guess"`
}

// GuessResponse represents the result of a Wordle guess.
type GuessResponse struct {
	Pattern    logic.StringPattern `json:"pattern"`
	IsAllMatch bool                `json:"isAllMatch"`
}

// GetGameWordHandler returns a random word for a new game.
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

// CheckGuessHandler processes a Wordle guess and returns the color pattern result.
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

// GetWordLengthsHandler returns a list of all word lengths available for play.
func GetWordLengthsHandler(c *fiber.Ctx) error {
	lengths := logic.GetWordLengths()
	return c.JSON(lengths)
}
