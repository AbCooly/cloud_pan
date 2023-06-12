package routes

import (
	"awesomeProject/app/controllers/app"
	"awesomeProject/app/middleware"
	"awesomeProject/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/register", app.Register)
	router.POST("/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.GET("/user/detail", app.Info)
		authRouter.GET("/user/file/list", app.GetUserFile)
		authRouter.POST("/user/dir/create", app.CreatDir)
		authRouter.POST("/user/repository/link", app.LinkRepository)
		authRouter.POST("/file/upload", app.PutFileUpload)
		authRouter.DELETE("user/file/delete", app.DeleteFile)
		authRouter.POST("/logout", app.Logout)
		//authRouter.POST("/image_upload", common.ImageUpload)

	}

}
