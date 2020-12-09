package controller

import (
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"../auth"
	"../crypt"
	"../db"
	"../model"
)

type UserController struct{}

type LoginResponse struct {
	Token string `json:"token"`
	Admin bool   `json:"admin"`
}

type CheckResponse struct {
	Exist bool `json:"exist"`
}

type RefreshResponse struct {
	Token string `json:"token"`
}

type UserProgressResponse struct {
	QuestionId []int `json:"progress"`
}

func (pc UserController) List(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	user := []model.User{}

	db := db.GormConnect()
	db.Find(&user)

	c.JSON(http.StatusOK, user)
}

func (pc UserController) Create(c *gin.Context) {
	user := model.User{}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	user.Password, err = crypt.PasswordEncrypt(user.Password)

	db := db.GormConnect()
	db.Create(&user)
	if db.NewRecord(user) == false {
		c.JSON(http.StatusOK, user)
	}
}

func (pc UserController) Login(c *gin.Context) {
	user := model.User{}
	users := model.User{}

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "request failed(json error)")
	}

	db := db.GormConnect()
	db.First(&users, "user_id=?", user.User_Id)

	err = crypt.CompareHashAndPassword(users.Password, user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, "Request is failed")
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user":     user.User_Id,
		"username": users.User_Name,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "secret"
	TokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.String(http.StatusInternalServerError, "Request is failed")
	}

	response := LoginResponse{
		TokenString,
		users.Admin,
	}

	c.JSON(http.StatusOK, response)
}

func (pc UserController) Delete(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	claims := token.Claims.(jwt.MapClaims)

	db := db.GormConnect()
	user := model.User{}
	db.Delete(&user, "user_id=?", claims["user"])

	c.String(http.StatusCreated, "completed user delete")
}

func (pc UserController) Edit(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	claims := token.Claims.(jwt.MapClaims)

	user := model.User{}
	user_json := model.User{}

	err = c.BindJSON(&user_json)
	if err != nil {
		c.String(http.StatusBadRequest, "request failed(json error)")
	}

	db := db.GormConnect()
	user.User_Id = claims["user"].(string)
	users_after := user
	db.First(&users_after)
	users_after.User_Name = user_json.User_Name
	db.Model(&user).Update(&users_after)
	db.Save(&users_after)

	c.String(http.StatusCreated, "completed username edit")
}

func (pc UserController) Check(c *gin.Context) {
	db := db.GormConnect()
	user := model.User{}

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "request failed(json error)")
	}

	if err := db.Where("user_id = ?", user.User_Id).First(&user).Error; err != nil {
		response := CheckResponse{
			false,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	response := CheckResponse{
		true,
	}

	c.JSON(http.StatusOK, response)
}

func (pc UserController) IdEdit(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	claims := token.Claims.(jwt.MapClaims)

	db := db.GormConnect()
	user := model.User{}
	user_json := model.User{}

	err = c.BindJSON(&user_json)
	if err != nil {
		c.String(http.StatusBadRequest, "request failed(json error)")
	}

	user.User_Id = claims["user"].(string)
	users_after := user
	db.First(&users_after)
	users_after.User_Id = user_json.User_Id
	db.Model(&user).Update(&users_after)
	db.Save(&users_after)
	c.String(http.StatusCreated, "completed userid edit")
}

func (pc UserController) PasswordEdit(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	claims := token.Claims.(jwt.MapClaims)

	db := db.GormConnect()
	user := model.User{}
	user_json := model.User{}

	err = c.BindJSON(&user_json)
	if err != nil {
		c.String(http.StatusBadRequest, "request failed(json error)")
	}

	user.User_Id = claims["user"].(string)
	users_after := user
	db.First(&users_after)
	users_after.Password, err = crypt.PasswordEncrypt(user_json.Password)
	db.Model(&user).Update(&users_after)
	db.Save(&users_after)
	c.String(http.StatusCreated, "complete password edit")
}

func (pc UserController) Refresh(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	claims := token.Claims.(jwt.MapClaims)

	db := db.GormConnect()
	user := model.User{}
	users := model.User{}
	db.First(&users, "user_id=?", claims["user"])

	token = jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user":     user.User_Id,
		"username": users.User_Name,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "secret"
	TokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.String(http.StatusInternalServerError, "Request is failed")
	}

	response := RefreshResponse{
		TokenString,
	}

	c.JSON(http.StatusOK, response)
}

func (pc UserController) TokenCheck(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "token is invalid")
		return
	}
	c.String(http.StatusOK, "token status ok")
}

func (pc UserController) UserProgress(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "token is invalid")
		return
	}
	claims := token.Claims.(jwt.MapClaims)

	db := db.GormConnect()
	user := model.User{}
	db.First(&user, "user_id=?", claims["user"])
	progress := []model.Progress{}
	db.Find(&progress, "user_id=?", user.ID)

	var proid []int
	for _, h := range progress {
		proid = append(proid, h.Pro_Id)
	}

	response := UserProgressResponse{
		proid,
	}

	c.JSON(http.StatusOK, response)
}
