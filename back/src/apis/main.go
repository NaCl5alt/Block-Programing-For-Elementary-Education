package main

import (
	"./db"
	"./model"
	"./router"
)

func main() {
	db := db.GormConnect()
	r := router.SetRouter(db)

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Problem{})
	db.AutoMigrate(&model.Progress{})
	db.AutoMigrate(&model.Hint{})

	defer db.Close()
	db.LogMode(true)

	r.Run(":80")
}
