package bootstrap

import (
	"awesomeProject/app/middleware"
	"awesomeProject/global"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	//router := gin.Default()

	if global.App.Config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	// 使用中间件
	router.Use(gin.Logger(), middleware.CustomRecovery())

	// 跨域处理
	// router.Use(middleware.Cors())

	// 前端项目静态资源
	router.StaticFile("/", "./static/dist/index.html")
	router.Static("/assets", "./static/dist/assets")
	router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")

	// 其他静态资源
	router.Static("/public", "./static")
	router.Static("/storage", "./storage/app/public")
	// 注册 api 分组路由
	apiGroup := router.Group("")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	err := r.Run(":" + global.App.Config.App.Port)
	if err != nil {
		global.App.Log.Error(err.Error())
		return
	}
}
