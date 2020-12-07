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
	api := r.Group("/api")
	{
		api.GET("/user", ctrl.List)
		api.POST("/user", ctrl.Create)
		api.POST("/login", ctrl.Login)
		api.DELETE("/user", ctrl.Delete)
		api.POST("/user/edit", ctrl.Edit)
		api.POST("/user/check", ctrl.Check)
		api.POST("/user/id", ctrl.IdEdit)
		api.POST("/user/password", ctrl.PasswordEdit)
		api.GET("/token", ctrl.Refresh)
		api.POST("/token", ctrl.TokenCheck)
		api.GET("/progress", ctrl.UserProgress)
		question := controller.QuestionController{}
		api.GET("/question", question.Get)
		api.POST("/question/:id", question.Answer)
		api.GET("/question/count", question.CountGet)
		api.GET("/question/paging", question.PagingGet)
	}

	return r
}
