package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"../auth"
	"../db"
	"../model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type QuestionController struct{}

type QuestionResponse struct {
<<<<<<< Updated upstream
	Id    int    `json:"qid"`
	Title string `json:"title"`
	Progress bool `json:"progress"`
=======
	Id       int    `json:"qid"`
	Title    string `json:"title"`
	Progress bool   `json:"progress"`
>>>>>>> Stashed changes
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

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	db := db.GormConnect()
<<<<<<< Updated upstream
	
=======

>>>>>>> Stashed changes
	problem := model.Problem{}
	db.First(&problem)

	user := model.User{}
	db.First(&user, "user_id=?", claims["user"])

	progress := []model.Progress{}
	db.Find(&progress, "user_id=?", user.ID)
<<<<<<< Updated upstream
	
=======

>>>>>>> Stashed changes
	count := 0
	db.Where("user_id=? AND pro_id=?", user.ID, problem.ID).Find(&progress).Count(&count)

	match := false
	if count == 0 {
		match = false
	} else {
		match = true
	}

	adf := QuestionResponse{
<<<<<<< Updated upstream
		Id:    int(problem.ID),
		Title: problem.Pro_Title,
=======
		Id:       int(problem.ID),
		Title:    problem.Pro_Title,
>>>>>>> Stashed changes
		Progress: match,
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

	token, err := auth.VerifyToken(tokenString)
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

		claims := token.Claims.(jwt.MapClaims)

		user := model.User{}
		user.User_Id = claims["user"].(string)
		db.First(&user)

		progress := model.Progress{}
		progress.Pro_Id = int(problem.ID)
		progress.User_Id = int(user.ID)

		if db.First(&progress).RecordNotFound() {
			db.Create(&progress)
		}

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
