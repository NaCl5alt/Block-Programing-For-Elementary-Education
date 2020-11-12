package server

import (
	"net/http"

	"../model"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		users := []model.User{}
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		user := model.User{}
		now := time.Now()
		user.CreatedAt = now
		user.UpdatedAt = now

		err := c.BindJSON(&user)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		db.NewRecord(user)
		db.Create(&user)
		db.LogMode(true)
		if db.NewRecord(user) == false {
			c.JSON(http.StatusOK, user)
		}
	})

	return r
}
