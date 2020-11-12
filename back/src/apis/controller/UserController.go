package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"../db"
	"../model"
)

type UserController struct{}

func (pc UserController) List(c *gin.Context) {
	db := db.GormConnect()
	users := []model.User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (pc UserController) Create(c *gin.Context) {
	db := db.GormConnect()
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
	if db.NewRecord(user) == false {
		c.JSON(http.StatusOK, user)
	}
}
