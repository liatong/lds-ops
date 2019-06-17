package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liatong/lds-ops/handler"

)


func TestMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		// Set example variable
		c.Set("example", "12345")
		//c.Set("db",&pool)
		// before request

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(TestMiddle())

	r.StaticFS("/download",http.Dir("/tmp/upload"))
	r.Static("/file","./file")
	r.Static("/static","./static")
	r.LoadHTMLGlob("templates/*")

	// Ping test
	r.GET("/ping",handler.Pong)

	r.POST("/upload",handler.UploadFile)
	r.POST("/package",handler.LishPackage)
	//r.POST("/test/UploadFile",handler.TestUploadFile)

 	//处理数据脚本
 	r.POST("/dbscript",handler.UploadScript)
 	/*
 	r.POST("/dbscriptlist",handler.ScriptList)
 	r.POST("/dbscript/exechistory",handler.DbExecHistory)
	*/


	r.GET("/packagelist",handler.PackageHtml)
	r.GET("/index",handler.IndexHtml)
	r.GET("/code",handler.CodeHtml)
	r.GET("dbupload",handler.DbuploadHtml)
	

	return r
}