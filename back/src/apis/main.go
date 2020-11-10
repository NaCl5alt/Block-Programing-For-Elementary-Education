package main

import (
	"./db"
	"./model"
	"./server"
)

func main() {
	db := db.GormConnect()
	r := server.SetRouter(db)

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&model.User{})

	defer db.Close()
	db.LogMode(true)

	r.Run(":80")
}
