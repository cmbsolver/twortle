package logic

import (
	"strings"
	"twortle/pkg/db"
	"twortle/pkg/db/tables"
)

func SearchWordleWords(text, contains, exclude []string) ([]tables.Word, error) {
	dbConn, _ := db.InitSQLiteConnection()
	words := tables.GetWordsByLength(dbConn, 5)
	db.CloseConnection(dbConn)
	return ParseWords(text, contains, exclude, words)
}

// ParseWords filters a list of WordEntry objects based on inclusion, exclusion, and positional criteria in the input parameters.
func ParseWords(text, contains, exclude []string, words []tables.Word) ([]tables.Word, error) {
	var preFilteredWords []tables.Word
	var filteredWords []tables.Word

	for _, word := range words {
		include := true

		if len(contains) > 0 {
			for _, letter := range contains {
				if !strings.Contains(word.WordText, letter) {
					include = false
					break
				}
			}
		}

		if len(exclude) > 0 {
			for _, letter := range exclude {
				if strings.Contains(word.WordText, letter) {
					include = false
					break
				}
			}
		}

		if include {
			preFilteredWords = append(preFilteredWords, word)
		}
	}

	if len(text) == 0 {
		return preFilteredWords, nil
	}

	for _, word := range preFilteredWords {
		include := true
		wordArray := strings.Split(word.WordText, "")
		for position, letter := range text {
			if letter != "%" {
				if letter != wordArray[position] {
					include = false
				}
			}
		}

		if include {
			filteredWords = append(filteredWords, word)
		}
	}

	return filteredWords, nil
}
