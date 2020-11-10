package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Age      int       `json:age`
	Birthday time.Time `json:birthday`
}
