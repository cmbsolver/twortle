package api

import (
	"twortle/pkg/logic"

	"github.com/gofiber/fiber/v2"
)

type PatternRow struct {
	PatternRow   []string `json:"patternRow"`
	PatternWords []string `json:"patternWords"`
}

type DrawRequest struct {
	PatternRow  []PatternRow `json:"patternRows"`
	PatternWord string       `json:"patternWord"`
}

type DrawResponse struct {
	PatternRows []PatternRow `json:"patternRows"`
}

func DrawPatternHandler(c *fiber.Ctx) error {
	var req DrawRequest
	var response DrawResponse
	response.PatternRows = make([]PatternRow, 0)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	//fmt.Println(req)

	patterns := make([]logic.Pattern, 0)
	for _, row := range req.PatternRow {
		tmpPattern := logic.BuildPatternFromStringArray(row.PatternRow)
		//fmt.Printf("%v - %v\n", row.PatternRow, tmpPattern)
		patterns = append(patterns, logic.Pattern{
			Length: len(tmpPattern),
			Colors: tmpPattern,
		})
	}

	for i, pattern := range patterns {
		words := logic.GetWordsFromPattern(req.PatternWord, pattern.Colors)
		response.PatternRows = append(response.PatternRows, PatternRow{
			PatternRow:   req.PatternRow[i].PatternRow,
			PatternWords: words,
		})
	}

	return c.JSON(response)
}
