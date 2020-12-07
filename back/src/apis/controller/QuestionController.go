package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"../auth"
	"../db"
	"../model"
	"github.com/gin-gonic/gin"
)

type QuestionController struct{}

type QuestionResponse struct {
	Id    int    `json:"qid"`
	Title string `json:"title"`
}
type CountResponse struct {
	Count int `json:"count"`
}

type QuestionHint struct {
	Hint string `json:"hint"`
}

type QuestionContents struct {
	Id      int            `json:"qid"`
	Title   string         `json:"title"`
	Content string         `json:"content"`
	Hints   []QuestionHint `json:"hints"`
}

type QuestionAnswer struct {
	IsCorrect bool `json:"accept"`
}

type AnswerJson struct {
	Answer string `json:"answer"`
}

type Json struct {
	Question_Id int `json:"qid"`
}

func (pc QuestionController) Get(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	db := db.GormConnect()
	problem := model.Problem{}

	db.First(&problem)
	c.String(http.StatusCreated, "complete edit")

	fmt.Println(problem)

	adf := QuestionResponse{
		Id:    int(problem.ID),
		Title: problem.Pro_Title,
	}

	c.JSON(http.StatusOK, adf)

}

func (pc QuestionController) Contents(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	db := db.GormConnect()
	problem := model.Problem{}
	hints := []model.Hint{}

	u64, err := strconv.ParseUint(c.Param("id"), 10, 0)
	problem.ID = uint(u64)

	db.First(&problem)

	db.Where(model.Hint{
		Pro_Id: int(problem.ID),
	}).Find(&hints)

	res_hints := []QuestionHint{}

	for _, h := range hints {
		res_hints = append(res_hints, QuestionHint{
			Hint: h.Hint,
		})
	}

	fmt.Println("qid:" + c.Param("id"))

	adf := QuestionContents{
		Id:      int(problem.ID),
		Title:   problem.Pro_Title,
		Content: problem.Pro_Content,
		Hints:   res_hints,
	}

	c.JSON(http.StatusOK, adf)
}

func (pc QuestionController) Answer(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	db := db.GormConnect()
	problem := model.Problem{}

	u64, err := strconv.ParseUint(c.Param("id"), 10, 0)
	problem.ID = uint(u64)

	db.First(&problem)

	json := AnswerJson{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accept := false
	if problem.Pro_Answer == json.Answer {
		accept = true
	} else {
		accept = false
	}

	adf := QuestionAnswer{
		IsCorrect: accept,
	}

	c.JSON(http.StatusOK, adf)
}

func (pc QuestionController) CountGet(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	var count int

	db := db.GormConnect()
	problems := model.Problem{}
	db.Model(&problems).Count(&count)

	adf := CountResponse{
		count,
	}

	c.JSON(http.StatusOK, adf)
}

func (pc QuestionController) PagingGet(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	json := Json{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := db.GormConnect()
	problem := []model.Problem{}

	qid := 1
	db.Limit(50).Where("ID >= ?", qid).Find(&problem)

	c.JSON(http.StatusOK, problem)
}
