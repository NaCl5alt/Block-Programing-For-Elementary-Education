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