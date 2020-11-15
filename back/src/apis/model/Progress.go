package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Progress struct {
	gorm.Model
	Pro_ID   int       `json:"proid"`
	User_ID  int       `json:"userid"`
	Rev_Time time.Time `json:"revtime"`
}
