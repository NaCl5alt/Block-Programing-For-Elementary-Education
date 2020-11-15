package model

import (
	"github.com/jinzhu/gorm"
)

type Problem struct {
	gorm.Model
	Pro_Title   string `json:"title"`
	Pro_Content string `json:"content"`
	Pro_Answer  string `json:"answer"`
	Hint_ID     int    `json:"hintid"`
}
