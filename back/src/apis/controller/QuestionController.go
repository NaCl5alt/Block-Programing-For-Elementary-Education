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

// func (pc QuestionController) countGET(c *gin.Context) {

// 	tokenString := c.Request.Header.Get("Authorization")
// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	token, err := auth.VerifyToken(tokenString)
// 	if err != nil {
// 		c.String(http.StatusUnauthorized, "Unauthorized")
// 		return
// 	}
// 	claims := token.Claims.(jwt.MapClaims)

// 	db := db.GormConnect()
// 	problems := model.Problem{}

// 	db.First(&problem)
// 	c.String(http.StatusCreated, "complete edit")
// 	adf := QuestionResponse{
// 		problem.count,
// 	}
// 	c.JSON(http.StatusOK,adf)
// }

// func (pc QuestionController) pagingGET(c *gin.Context) {
// 	type QuestionResponse struct {
// 		Id int `json:"qid"`
// 		Title string `json:"title"`
// 	}

// 	tokenString := c.Request.Header.Get("Authorization")
// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	token, err := auth.VerifyToken(tokenString)
// 	if err != nil {
// 		c.String(http.StatusUnauthorized, "Unauthorized")
// 		return
// 	}
// 	claims := token.Claims.(jwt.MapClaims)

// 	db := db.GormConnect()
// 	problem := model.Problem{}

// 	for i := 0; i < 50; i++{
// 		db.First(&problem)
// 		c.String(http.StatusCreated, "complete edit")
// 		adf := QuestionResponse{
// 			problem,
// 			problem.pro_title,
// 		}
// 		c.JSON(http.StatusOK,adf)
// 	}
// }
