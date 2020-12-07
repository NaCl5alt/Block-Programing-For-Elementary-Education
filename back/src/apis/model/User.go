package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	User_Id    string `json:"userid"`
	User_Name  string `json:"name"`
	Password   string `json:"password"`
	Admin      bool   `json:"admin"`
	DeleteFlag bool   `json:"delete"`
}
