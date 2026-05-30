package tables

import (
	"bufio"
	"os"
	"slices"
	"strings"

	"gorm.io/gorm"
)

// Word represents a word entry in the database for the Wordle game.
type Word struct {
	gorm.Model
	WordText   string
	WordLength int
	Played     bool
}

// TableName returns the table name for the Word model.
func (Word) TableName() string {
	return "words"
}

// GetWordCount returns the total number of words in the database.
func GetWordCount(db *gorm.DB) int64 {
	var count int64
	db.Model(&Word{}).Count(&count)
	return count
}

// GetRandomWordByLength returns a random word of the specified length that hasn't been played yet.
func GetRandomWordByLength(db *gorm.DB, length int) Word {
	var word Word
	db.
		Where("word_length = ?", length).
		Where("played = ?", false).
		Order("RANDOM()").First(&word)

	return word
}

// GetWordsByLength returns all words of the specified length.
func GetWordsByLength(db *gorm.DB, length int) []Word {
	var words []Word
	db.Where("word_length = ?", length).Find(&words)
	return words
}

// GetAllWords returns all words in the database.
func GetAllWords(db *gorm.DB) []Word {
	var words []Word
	db.Find(&words)

	return words
}

// GetWordLengths returns a list of unique word lengths present in the database.
func GetWordLengths(db *gorm.DB) []int {
	var lengths []int

	for _, word := range GetAllWords(db) {
		if slices.Contains(lengths, word.WordLength) == false {
			lengths = append(lengths, word.WordLength)
		}
	}

	return lengths
}

// AddWords inserts multiple word entries into the database.
func AddWords(db *gorm.DB, words []Word) {
	db.Create(&words)
}

// LoadFile reads words from a text file and loads them into the database.
func LoadFile(db *gorm.DB, filename string) {
	words := make([]Word, 0)
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			upperWord := strings.ToUpper(word)

			if IsAllLetter(upperWord) {
				words = append(words,
					Word{
						WordText:   upperWord,
						WordLength: len(strings.Split(upperWord, "")),
						Played:     false,
					})
			}

			if len(words) >= 750 {
				AddWords(db, words)
				words = make([]Word, 0)
			}
		}
	}

	if len(words) > 0 {
		AddWords(db, words)
	}
}

// IsAllLetter checks if a string contains only English letters (A-Z).
func IsAllLetter(word string) bool {
	alphaArray := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	wordArray := strings.Split(word, "")

	for _, letter := range wordArray {
		if !slices.Contains(alphaArray, letter) {
			return false
		}
	}
	return true
}

// UpdatePlayed marks a word as played in the database by its ID.
func UpdatePlayed(db *gorm.DB, wordId uint) {
	db.Model(&Word{}).Where("id = ?", wordId).Update("played", true)
}
