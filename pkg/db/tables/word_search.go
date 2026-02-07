package tables

import (
	"gorm.io/gorm"
)

type WordSearch struct {
	gorm.Model
	Grid       string `gorm:"type:text"` // JSON representation of the grid
	Words      string `gorm:"type:text"` // JSON representation of the words to find
	FoundWords string `gorm:"type:text"` // JSON representation of the words already found
	WordCount  int
	GridSize   int
}

func (WordSearch) TableName() string {
	return "word_searches"
}
