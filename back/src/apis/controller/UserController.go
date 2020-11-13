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

func (pc UserController) List(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	db := db.GormConnect()
	users := []model.User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (pc UserController) Create(c *gin.Context) {
	db := db.GormConnect()
	user := model.User{}
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	user.Password, err = crypt.PasswordEncrypt(user.Password)
	db.NewRecord(user)
	db.Create(&user)
	if db.NewRecord(user) == false {
		c.JSON(http.StatusOK, user)
	}
}

func (pc UserController) Login(c *gin.Context) {
	db := db.GormConnect()
	user := model.User{}
	users := model.User{}

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "request failed(json error)")
	}
	db.First(&users, "user_id=?", user.User_ID)

	err = crypt.CompareHashAndPassword(users.Password, user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, "Request is failed")
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user":     user.User_ID,
		"username": users.User_Name,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "secret"
	TokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.String(http.StatusInternalServerError, "Request is failed")
	}

	adf := LoginResponse{
		TokenString,
		users.Admin,
	}

	c.JSON(http.StatusOK, adf)
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
	users := model.User{}
	db.Delete(&users, "user_id=?", claims["user"])
	c.String(http.StatusOK, "complete delete")
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

	db := db.GormConnect()
	users := model.User{}
	user := model.User{}

	err = c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "request failed(json error)")
	}

	users.User_ID = claims["user"].(string)
	users_after := users
	db.First(&users_after)
	users_after.User_Name = user.User_Name
	db.Model(&users).Update(&users_after)
	db.Save(&users_after)
	c.String(http.StatusOK, "complete edit")
}
