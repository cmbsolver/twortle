package logic

import (
	"twortle/pkg/db"
	"twortle/pkg/db/tables"
)

func GetWordForGame(length int) tables.Word {
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)
	word := tables.GetRandomWordByLength(dbConn, length)
	return word
}

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

func GetWordLengths() []int {
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)
	return tables.GetWordLengths(dbConn)
}
