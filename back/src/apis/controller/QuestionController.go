package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	// jwt "github.com/dgrijalva/jwt-go"
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

type QuestionAnswer struct {
	isCorrect bool `json:"accept"`
}

func (pc QuestionController) Get(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	// _,ok := token.Claims.(jwt.MapClaims)

	// if !ok {
	// 	panic(ok)
	// }

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

	fmt.Println("qid:" + c.Param("id"))

	adf := QuestionAnswer{
		isCorrect: problem.Pro_Answer == c.PostForm("answer"),
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

func (pc QuestionController) pagingGET(c *gin.Context) {
	type QuestionResponse struct {
		Id    int    `json:"qid"`
		Title string `json:"title"`
	}
	type JsonRequest struct {
		FieldStr  string `json:"field_str"`
		FieldInt  int    `json:"field_int"`
		FieldBool bool   `json:"field_bool"`
	}
	type Json struct {
		Question_Id int `json:"qid"`
	}
	json := Json{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	// claims := token.Claims.(jwt.MapClaims)

	db := db.GormConnect()
	problem := model.Problem{}

	qid := 10
	db.Limit(50).Where("qid > ?", qid).Find(&problem)
	c.String(http.StatusCreated, "complete edit")
	adf := QuestionResponse{
		Id:    int(problem.ID),
		Title: problem.Pro_Title,
	}
	// adf := QuestionResponse{
	// 	Id: int(problem.ID),
	// 	Title: problem.Pro_Title,
	// }
	c.JSON(http.StatusOK, adf)
}
