package controller

import (
	"net/http"
	"strconv"
	"strings"
	"fmt"
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

type QuestionHintDetail struct {
	Hint string `json:"hint"`
}

type QuestionDetail struct {
	Title   string         `json:"title"`
	Content string         `json:"content"`
	Answer string `json:"answer"`
	Hints   []QuestionHint `json:"hints"`
}

// type AnswerJson struct {
	// Answer string `json:"answer"`
// }

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

func (pc AdminController) DetailGET(c *gin.Context) {
	/*
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	fmt.Println("debug")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
*/
	db := db.GormConnect()
	problem := model.Problem{}
	hints := []model.Hint{}

	fmt.Println("debug")


	u64, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	problem.ID = uint(u64)

	db.First(&problem)

	db.Where(model.Hint{
		Pro_Id: int(problem.ID),
	}).Find(&hints)
	
	json := AnswerJson{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res_hints := []QuestionHint{}

	for _, h := range hints {
		res_hints = append(res_hints, QuestionHint{
			Hint: h.Hint,
		})
	}

	fmt.Println("qid:" + c.Param("id"))

	adf := QuestionDetail{
		Title:   problem.Pro_Title,
		Content: problem.Pro_Content,
		Answer: problem.Pro_Answer,
		Hints:   res_hints,
	}

	c.JSON(http.StatusOK, adf)
}


