package logic

import (
	"twortle/pkg/db"
	"twortle/pkg/db/tables"
)

// GetWordForGame retrieves a random word of the specified length from the database.
func GetWordForGame(length int) tables.Word {
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)
	word := tables.GetRandomWordByLength(dbConn, length)
	return word
}

// CheckWordPattern compares a guess against the target word and returns the color pattern.
// If the guess is correct, it updates the word's played status in the database.
func CheckWordPattern(guess string, word tables.Word) (StringPattern, bool) {
	isAllMatch := true
	result, _ := GetColorPatternFromWords(word.WordText, guess)
	for _, color := range result.Colors {
		if color == Grey || color == Yellow {
			isAllMatch = false
			break
		}
	}

	if isAllMatch {
		dbConn, _ := db.InitSQLiteConnection()
		defer db.CloseConnection(dbConn)
		tables.UpdatePlayed(dbConn, word.ID)
	}

	return BuildStringPatternFromPattern(result.Colors), isAllMatch
}

// GetWordLengths returns a list of all unique word lengths available in the database.
func GetWordLengths() []int {
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)
	return tables.GetWordLengths(dbConn)
}
