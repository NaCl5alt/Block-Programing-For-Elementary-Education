package router

import (
	"../controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	user := controller.UserController{}
	question := controller.QuestionController{}
	admin := controller.AdminController{}
	api := r.Group("/api")
	{
		api.GET("/user", user.List)
		api.POST("/user", user.Create)
		api.POST("/login", user.Login)
		api.DELETE("/user", user.Delete)
		api.POST("/user/edit", user.Edit)
		api.POST("/user/check", user.Check)
		api.POST("/user/id", user.IdEdit)
		api.POST("/user/password", user.PasswordEdit)
		api.GET("/token", user.Refresh)
		api.POST("/token", user.TokenCheck)
		api.GET("/progress", user.UserProgress)

		api.GET("/question", question.GetFirst)
		api.GET("/question/contents/:id", question.Contents)
		api.POST("/question/answer/:id", question.Answer)
		api.GET("/question/count", question.CountGet)
		api.GET("/question/paging", question.PagingGet)

		api.POST("/question", admin.AddQuestion)
		api.POST("/question/edit/:id", admin.EditQuestion)
		api.GET("/user/progress/:id", admin.UserIdProgress)
		api.GET("/user/progress", admin.AllProgress)
		api.GET("/question/detail/:id", admin.DetailGET)
		api.DELETE("/question/:id", admin.Delete)
	}
	return r
}
