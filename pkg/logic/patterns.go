package logic

import (
	"errors"
	"slices"
	"strings"
	"twortle/pkg/db"
	"twortle/pkg/db/tables"
)

type Color int

const (
	Grey Color = iota
	Yellow
	Green
)

type Pattern struct {
	Length int
	Colors []Color
}

type StringPattern struct {
	Length int
	Colors []string
}

func BuildStringPatternFromPattern(pattern []Color) StringPattern {
	stringPattern := make([]string, 0)
	for i, color := range pattern {
		switch color {
		case Grey:
			stringPattern[i] = "grey"
			break
		case Yellow:
			stringPattern[i] = "yellow"
			break
		case Green:
			stringPattern[i] = "green"
			break
		}
	}

	return StringPattern{Length: len(pattern), Colors: stringPattern}
}

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

func GetColorPatternFromWords(patternWord, word string) (Pattern, error) {
	patternArray := strings.Split(patternWord, "")
	wordArray := strings.Split(word, "")
	return GetColorPatternFromArrays(patternArray, wordArray)
}

func GetColorPatternFromArrays(patternWord, word []string) (Pattern, error) {
	if len(patternWord) != len(word) {
		return Pattern{}, errors.New("pattern and word must be of equal length")
	}

	colorPattern := make([]Color, len(patternWord))

	for i, letter := range word {
		if slices.Contains(patternWord, letter) {
			colorPattern[i] = Yellow
		} else {
			colorPattern[i] = Grey
		}
	}

	for i, letter := range word {
		if patternWord[i] == letter {
			colorPattern[i] = Green
		}
	}

	return Pattern{Length: len(patternWord), Colors: colorPattern}, nil
}
