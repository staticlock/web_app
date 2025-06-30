package router

import (
	"log"
	"time"
	"web_app/controllers"
	"web_app/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {
	// 设置 zap 日志输出
	// 1. 完全接管标准库log输出
	log.SetOutput(logger.GetGinWriter())
	log.SetFlags(0) // 禁用标准库的前缀

	// 2. 重定向Gin的所有输出
	gin.DisableConsoleColor()                      // 禁用Gin默认的日志输出
	gin.DefaultWriter = logger.GetGinWriter()      // 重定向Gin的普通日志
	gin.DefaultErrorWriter = logger.GetGinWriter() // 重定向Gin的错误日志
	gin.SetMode(gin.ReleaseMode)                   //设置为生产环境，减少日志输出
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(cors.New(cors.Config{
		//前端地址
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	apiV1 := r.Group("/api/v1")
	//括号可以省略，为了区分路由，留下
	{
		//路径参数 GET /get/123
		apiV1.GET("/get/a/:id", controllers.TestFunc1)
		//查询参数 GET /get?page=2&size=10
		apiV1.GET("/get", controllers.TestFunc2)
		//文件下载 GET /get/download/web_app.log
		apiV1.GET("/get/download/:filename", controllers.Download)
	}
	{
		//路径参数 POST /post/123
		apiV1.POST("/post/a/:id", controllers.TestFunc1)
		//查询参数 POST /post?page=2&size=10
		apiV1.POST("/post", controllers.TestFunc2)
		//JSON参数 POST
		apiV1.POST("/post/json", controllers.TestFunc3)
		//表单参数 POST
		apiV1.POST("/post/form", controllers.TestFunc4)
		//单文件上传
		apiV1.POST("/post/upload", controllers.TestFunc5)
		//多文件上传
		apiV1.POST("/post/uploadFiles", controllers.TestMultiFileUpload)
		//二进制文件
		apiV1.POST("/post/bin", controllers.TestFunc6)
	}
	{
		//路径参数 PUT /put/123
		apiV1.PUT("/put/a/:id", controllers.TestFunc4)
		//JSON参数 PUT
		apiV1.PUT("/put/json", controllers.TestFunc3)
	}
	{
		//路径参数 DELETE  /delete/123
		apiV1.DELETE("/delete/:id", controllers.TestFunc4)
	}
	apiV2 := r.Group("/api/v2")
	apiV2.GET("/getExchangeRates", controllers.TestFunc4)
	{
		apiV2.POST("/createExchangeRate", controllers.TestFunc4)
		apiV2.POST("/articles", controllers.TestFunc4)
		apiV2.GET("/articles", controllers.TestFunc4)
		apiV2.GET("/articles/:id", controllers.TestFunc4)
		apiV2.POST("/articles/:id/like", controllers.TestFunc4)
		apiV2.GET("/articles/:id/like", controllers.TestFunc4)
	}
	return r
}
