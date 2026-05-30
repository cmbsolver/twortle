package logic

import (
	"errors"
	"slices"
	"strings"
	"twortle/pkg/db"
	"twortle/pkg/db/tables"
)

// Color represents the state of a letter in a Wordle guess (Grey, Yellow, Green).
type Color int

const (
	// Grey indicates the letter is not in the word.
	Grey Color = iota
	// Yellow indicates the letter is in the word but in the wrong position.
	Yellow
	// Green indicates the letter is in the word and in the correct position.
	Green
)

// Pattern represents a color sequence for a Wordle result.
type Pattern struct {
	Length int
	Colors []Color
}

// StringPattern is a JSON-friendly representation of a Pattern.
type StringPattern struct {
	Length int      `json:"length"`
	Colors []string `json:"colors"`
}

// BuildStringPatternFromPattern converts a Color slice into a StringPattern.
func BuildStringPatternFromPattern(pattern []Color) StringPattern {
	stringPattern := make([]string, 0)
	for _, color := range pattern {
		switch color {
		case Grey:
			stringPattern = append(stringPattern, "grey")
			break
		case Yellow:
			stringPattern = append(stringPattern, "yellow")
			break
		case Green:
			stringPattern = append(stringPattern, "green")
			break
		}
	}

	return StringPattern{Length: len(pattern), Colors: stringPattern}
}

// BuildPatternFromStringArray converts a slice of color names ("grey", "yellow", "green") into a Color slice.
func BuildPatternFromStringArray(stringPattern []string) []Color {
	colors := make([]Color, 0)
	for _, letter := range stringPattern {
		switch letter {
		case "grey":
			colors = append(colors, Grey)
			break
		case "yellow":
			colors = append(colors, Yellow)
			break
		case "green":
			colors = append(colors, Green)
		}
	}

	return colors
}

// GetWordsFromPattern returns all words in the database that match the given pattern when compared to patternWord.
func GetWordsFromPattern(patternWord string, colorPattern []Color) []string {
	words := make([]string, 0)
	dbConn, _ := db.InitSQLiteConnection()
	defer db.CloseConnection(dbConn)

	dictWords := tables.GetWordsByLength(dbConn, len(strings.Split(patternWord, "")))
	for _, word := range dictWords {
		wordPattern, _ := GetColorPatternFromWords(patternWord, word.WordText)
		//fmt.Printf("%v - %v\n", wordPattern, colorPattern)
		if slices.Equal(wordPattern.Colors, colorPattern) {
			words = append(words, word.WordText)
		}
	}

	return words
}

// GetColorPatternFromWords calculates the Wordle color pattern for a guess compared to a target word.
func GetColorPatternFromWords(patternWord, word string) (Pattern, error) {
	patternArray := strings.Split(patternWord, "")
	wordArray := strings.Split(word, "")
	return GetColorPatternFromArrays(patternArray, wordArray)
}

// GetColorPatternFromArrays calculates the Wordle color pattern for a guess compared to a target word, both as string slices.
func GetColorPatternFromArrays(patternWord, word []string) (Pattern, error) {
	if len(patternWord) != len(word) {
		return Pattern{}, errors.New("pattern and word must be of equal length")
	}

	colorPattern := make([]Color, len(patternWord))

	patternClone := slices.Clone(patternWord)

	for i, letter := range word {
		if patternClone[i] == letter {
			colorPattern[i] = Green
			patternClone[i] = "|"
		}
	}

	isAllGreen := true
	for _, color := range colorPattern {
		if color != Green {
			isAllGreen = false
		}
	}

	if isAllGreen {
		return Pattern{Length: len(patternWord), Colors: colorPattern}, nil
	}

	for i, letter := range word {
		if idx := slices.Index(patternClone, letter); idx != -1 {
			colorPattern[i] = Yellow

			// Remove the specific instance found at idx
			patternClone[idx] = "|"
		} else {
			if colorPattern[i] != Green {
				colorPattern[i] = Grey
			}
		}
	}

	return Pattern{Length: len(patternWord), Colors: colorPattern}, nil
}
