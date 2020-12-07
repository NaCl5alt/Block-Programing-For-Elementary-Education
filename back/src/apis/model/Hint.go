package model

import (
	"github.com/jinzhu/gorm"
)

type Hint struct {
	gorm.Model
	Pro_Id int    `json:"proid"`
	Hint   string `json:"hint"`
}
