package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"apis/model"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "user"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "BlockProgram"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
	db := gormConnect()

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&model.User{})

	defer db.Close()
	db.LogMode(true)
}
