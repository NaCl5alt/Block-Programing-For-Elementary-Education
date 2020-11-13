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
	r.GET("/user", ctrl.List)
	r.POST("/user", ctrl.Create)
	r.POST("/login", ctrl.Login)
	r.DELETE("/user", ctrl.Delete)
	r.POST("/user/edit", ctrl.Edit)
	r.POST("/user/check", ctrl.Check)
	r.POST("/user/id", ctrl.IdEdit)
	r.POST("/user/password", ctrl.PasswordEdit)
	r.GET("/token", ctrl.Refresh)
	r.POST("/token", ctrl.TokenCheck)

	return r
}
