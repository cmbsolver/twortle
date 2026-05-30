package logic

import (
	"math/rand/v2"
	"twortle/pkg/db"
	"twortle/pkg/db/tables"
)

// WordSearchGame represents the state and data for a word search game.
type WordSearchGame struct {
	Grid       [][]string `json:"grid"`
	Words      []string   `json:"words"`
	FoundWords []string   `json:"foundWords"`
}

// GenerateWordSearch creates a new word search game with the specified number of random words.
func GenerateWordSearch(wordCount int) WordSearchGame {
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)

	var allWords []tables.Word
	dbConn.Order("RANDOM()").Limit(wordCount).Find(&allWords)

	words := make([]string, 0, len(allWords))
	maxLength := 0
	for _, w := range allWords {
		words = append(words, w.WordText)
		if len(w.WordText) > maxLength {
			maxLength = len(w.WordText)
		}
	}

	gridSize := max(maxLength+2, wordCount+2, 10)
	grid := make([][]string, gridSize)
	for i := range gridSize {
		grid[i] = make([]string, gridSize)
	}

	directions := [][2]int{
		{0, 1},   // Horizontal forward
		{0, -1},  // Horizontal backward
		{1, 0},   // Vertical forward
		{-1, 0},  // Vertical backward
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{1, -1},  // Diagonal down-left
		{-1, 1},  // Diagonal up-right
	}

	for _, word := range words {
		placed := false
		attempts := 0
		for !placed && attempts < 100 {
			attempts++
			dir := directions[rand.IntN(len(directions))]
			startX := rand.IntN(gridSize)
			startY := rand.IntN(gridSize)

			if canPlace(grid, word, startX, startY, dir) {
				place(grid, word, startX, startY, dir)
				placed = true
			}
		}
	}

	fillEmpty(grid)

	return WordSearchGame{
		Grid:       grid,
		Words:      words,
		FoundWords: []string{},
	}
}

func canPlace(grid [][]string, word string, startX, startY int, dir [2]int) bool {
	gridSize := len(grid)
	for i := range len(word) {
		x := startX + i*dir[0]
		y := startY + i*dir[1]

		if x < 0 || x >= gridSize || y < 0 || y >= gridSize {
			return false
		}

		if grid[x][y] != "" && grid[x][y] != string(word[i]) {
			return false
		}
	}
	return true
}

func place(grid [][]string, word string, startX, startY int, dir [2]int) {
	for i := range len(word) {
		x := startX + i*dir[0]
		y := startY + i*dir[1]
		grid[x][y] = string(word[i])
	}
}

func fillEmpty(grid [][]string) {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "" {
				grid[i][j] = string(letters[rand.IntN(len(letters))])
			}
		}
	}
}
