package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"../auth"
	"../db"
	"../model"
)

type QuestionController struct{}

func (pc QuestionController) List(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	db := db.GormConnect()
	users := []model.Problem{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}
