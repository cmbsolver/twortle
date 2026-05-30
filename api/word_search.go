package api

import (
	"encoding/json"
	"twortle/pkg/db"
	"twortle/pkg/db/tables"
	"twortle/pkg/logic"

	"github.com/gofiber/fiber/v2"
)

// WordSearchGenerateRequest represents a request to generate a new word search game.
type WordSearchGenerateRequest struct {
	WordCount int `json:"wordCount"`
}

// WordSearchSaveRequest represents a request to save or update a word search game state.
type WordSearchSaveRequest struct {
	ID         uint       `json:"id"`
	Grid       [][]string `json:"grid"`
	Words      []string   `json:"words"`
	FoundWords []string   `json:"foundWords"`
	WordCount  int        `json:"wordCount"`
}

// GenerateWordSearchHandler handles requests to generate a new random word search grid.
func GenerateWordSearchHandler(c *fiber.Ctx) error {
	var req WordSearchGenerateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if req.WordCount <= 0 {
		req.WordCount = 10
	}

	game := logic.GenerateWordSearch(req.WordCount)
	return c.JSON(game)
}

// SaveWordSearchHandler handles requests to save a new or update an existing word search game.
func SaveWordSearchHandler(c *fiber.Ctx) error {
	var req WordSearchSaveRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	gridJSON, _ := json.Marshal(req.Grid)
	wordsJSON, _ := json.Marshal(req.Words)
	foundWordsJSON, _ := json.Marshal(req.FoundWords)

	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)

	if req.ID > 0 {
		// Update existing
		var ws tables.WordSearch
		if err := dbConn.First(&ws, req.ID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Word search not found",
			})
		}
		ws.Grid = string(gridJSON)
		ws.Words = string(wordsJSON)
		ws.FoundWords = string(foundWordsJSON)
		ws.WordCount = req.WordCount
		ws.GridSize = len(req.Grid)

		if err := dbConn.Save(&ws).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update word search",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Word search updated successfully",
			"id":      ws.ID,
		})
	}

	// Create new
	ws := tables.WordSearch{
		Grid:       string(gridJSON),
		Words:      string(wordsJSON),
		FoundWords: string(foundWordsJSON),
		WordCount:  req.WordCount,
		GridSize:   len(req.Grid),
	}

	if err := dbConn.Create(&ws).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save word search",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Word search saved successfully",
		"id":      ws.ID,
	})
}

// DeleteWordSearchHandler handles requests to delete a word search game by its ID.
func DeleteWordSearchHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)

	if err := dbConn.Delete(&tables.WordSearch{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete word search",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Word search deleted successfully",
	})
}

// ListWordSearchesHandler returns a list of all saved word search games.
func ListWordSearchesHandler(c *fiber.Ctx) error {
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)

	var searches []tables.WordSearch
	if err := dbConn.Select("id", "word_count", "grid_size", "created_at").Order("created_at desc").Find(&searches).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to list word searches",
		})
	}

	return c.JSON(searches)
}

// GetWordSearchHandler retrieves a single word search game by its ID.
func GetWordSearchHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)

	var ws tables.WordSearch
	if err := dbConn.First(&ws, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Word search not found",
		})
	}

	var grid [][]string
	var words []string
	var foundWords []string
	json.Unmarshal([]byte(ws.Grid), &grid)
	json.Unmarshal([]byte(ws.Words), &words)
	json.Unmarshal([]byte(ws.FoundWords), &foundWords)

	return c.JSON(fiber.Map{
		"id":         ws.ID,
		"grid":       grid,
		"words":      words,
		"foundWords": foundWords,
		"wordCount":  ws.WordCount,
		"gridSize":   ws.GridSize,
	})
}
