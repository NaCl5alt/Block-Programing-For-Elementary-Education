package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"../auth"
	"../db"
	"../model"
)

type AdminController struct{}

type ProgressResponse struct {
	AllProgress []ProgressUserId `json:"prog"`
}

type ProgressUserId struct {
	UserId     int   `json:"uid"`
	QuestionId []int `json:"progress"`
}

func (pc AdminController) AllProgress(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "token is invalid")
		return
	}

	db := db.GormConnect()
	user := []model.User{}
	prog := []ProgressUserId{}
	var count int
	db.Model(&user).Count(&count)
	db.Find(&user)

	for _, h := range user {
		progress := []model.Progress{}
		var cnt int
		db.Model(&progress).Where("user_id=?", h.ID).Count(&cnt)
		db.Where("user_id=?", h.ID).Find(&progress)
		var proid []int
		for j := 0; j < cnt; j++ {
			proid = append(proid, progress[j].Pro_Id)
		}

		prog = append(prog, ProgressUserId{
			UserId:     progress[0].User_Id,
			QuestionId: proid,
		})
	}
	adf := ProgressResponse{
		prog,
	}

	c.JSON(http.StatusOK, adf)
}

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
