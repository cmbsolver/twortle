package tables

import (
	"gorm.io/gorm"
)

// WordSearch represents a word search game state in the database.
type WordSearch struct {
	gorm.Model
	Grid       string `gorm:"type:text"` // JSON representation of the grid
	Words      string `gorm:"type:text"` // JSON representation of the words to find
	FoundWords string `gorm:"type:text"` // JSON representation of the words already found
	WordCount  int
	GridSize   int
}

// TableName returns the table name for the WordSearch model.
func (WordSearch) TableName() string {
	return "word_searches"
}
