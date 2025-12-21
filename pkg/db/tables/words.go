package tables

import (
	"bufio"
	"os"
	"strings"

	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	WordText   string
	WordLength int
}

func (Word) TableName() string {
	return "words"
}

func GetWordCount(db *gorm.DB) int64 {
	var count int64
	db.Model(&Word{}).Count(&count)
	return count
}

func GetRandomWord(db *gorm.DB) string {
	var word Word
	db.Order("RANDOM()").First(&word)
	return word.WordText
}

func GetAllWords(db *gorm.DB) []string {
	var words []Word
	db.Find(&words)
	var result []string
	for _, word := range words {
		result = append(result, word.WordText)
	}
	return result
}

func GetWordsByLength(db *gorm.DB, length int) []string {
	var words []Word
	db.Where("word_length = ?", length).Find(&words)
	var result []string
	for _, word := range words {
		result = append(result, word.WordText)
	}
	return result
}

func AddWord(db *gorm.DB, word string) {
	db.Create(&Word{WordText: word})
}

func AddWords(db *gorm.DB, words []Word) {
	db.Create(&words)
}

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
			words = append(words,
				Word{
					WordText:   upperWord,
					WordLength: len(strings.Split(upperWord, "")),
				})

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
