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

type AdminController struct{}

type LoginResponse struct {
	Token string `json:"token"`
	Admin bool   `json:"admin"`
}