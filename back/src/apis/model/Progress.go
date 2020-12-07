package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Progress struct {
	gorm.Model
	Pro_Id   int       `json:"proid"`
	User_Id  int       `json:"userid"`
	Rev_Time time.Time `json:"revtime"`
}
