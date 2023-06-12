package main

import (
	"awesomeProject/bootstrap"
	"awesomeProject/global"
)

func main() {
	// 初始化config
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化验证器
	bootstrap.InitializeValidator()

	//初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 初始化MinIO
	global.App.MinioClient = bootstrap.InitMinIO()

	// 初始化文件系统
	bootstrap.InitializeStorage()

	// 启动服务器
	bootstrap.RunServer()
}
