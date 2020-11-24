package model

import (
	"github.com/jinzhu/gorm"
)

type Problem struct {
	gorm.Model
	Hint_To_ID     int    `json:"hinttoid"`
	Hint_ID      int    `json:"hintid"`
}