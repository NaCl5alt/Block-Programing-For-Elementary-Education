package controller

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

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
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	db.First(&users, "user_id=?", user.User_ID)

	err = crypt.CompareHashAndPassword(users.Password, user.Password)
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user": user.User_ID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	var secretKey = "secret"
	TokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	adf := LoginResponse{
		TokenString,
		users.Admin,
	}

	c.JSON(http.StatusOK, adf)
}
