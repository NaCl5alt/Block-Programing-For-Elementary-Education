package controller

import (
	//"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"../auth"
	"../db"
	"../model"
)

type AdminController struct{}

func (pc AdminController) UserIdProgress(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "token is invalid")
		return
	}

	u64, err := strconv.ParseUint(c.Param("id"), 10, 0)
	userid := uint(u64)
	db := db.GormConnect()
	progress := []model.Progress{}
	db.Find(&progress, "user_id=?", userid)
	var count int
	db.Model(&progress).Where("user_id = ?", userid).Count(&count)

	var proid []int
	for i := 0; i < count; i++ {
		proid = append(proid, progress[i].Pro_Id)
	}

	type ResponseProgress struct {
		QuestionId []int `json:"progress"`
	}

	adf := ResponseProgress{
		proid,
	}

	c.JSON(http.StatusOK, adf)
}
