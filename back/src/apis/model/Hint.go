package model

import (
	"github.com/jinzhu/gorm"
)

type Hint struct {
	gorm.Model
	Hint_ID int    `json:"hintid"`
	Hint    string `json:"hint"`
}
