package server

import (
	"net/http"

	"../model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	//READ
	//全レコード
	r.GET("/users", func(c *gin.Context) {
		users := []model.User{}
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	return r
}
