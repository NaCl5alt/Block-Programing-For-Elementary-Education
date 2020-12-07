package controller

import (
	"net/http"
	"strings"
	"fmt"

	// jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"../auth"
	"../db"
	"../model"
)

type QuestionController struct{}
type QuestionResponse struct {
	Id int `json:"qid"`
	Title string `json:"title"`
}
type CountResponse struct{
	Count int `json:"count"`
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
		Id: int(problem.ID),
		Title: problem.Pro_Title,
	}

	c.JSON(http.StatusOK, adf)

}
func (pc QuestionController) countGET(c *gin.Context) {
	
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// _, err := auth.VerifyToken(tokenString)
	// if err != nil {
	// 	c.String(http.StatusUnauthorized, "Unauthorized")
	// 	return
	// }
	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	db := db.GormConnect()
	problems := model.Problem{}

	result := db.First(&problems)
	c.String(http.StatusCreated, "complete edit")
	adf := CountResponse{
		Count: int(result.RowsAffected),
	}
	c.JSON(http.StatusOK,adf)
}

func (pc QuestionController) pagingGET(c *gin.Context) {
	type QuestionResponse struct {
		Id int `json:"qid"`
		Title string `json:"title"`
	}
	type JsonRequest struct {
		FieldStr  string `json:"field_str"`
		FieldInt  int    `json:"field_int"`
		FieldBool bool   `json:"field_bool"`
	}
	type Json struct {
		Question_Id int  `json:"qid"`
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
		db.Limit(50).Where("qid > ?",qid).Find(&problem)
		c.String(http.StatusCreated, "complete edit")
		adf := QuestionResponse{
			Id: int(problem.ID),
			Title: problem.Pro_Title,
		}
		// adf := QuestionResponse{
		// 	Id: int(problem.ID),
		// 	Title: problem.Pro_Title,
		// }
		c.JSON(http.StatusOK,adf)
}
