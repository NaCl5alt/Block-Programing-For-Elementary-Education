package router

import (
	"../controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	ctrl := controller.UserController{}
	r.GET("/users", ctrl.List)
	r.POST("/users", ctrl.Create)
	r.POST("/login", ctrl.Login)

	return r
}
