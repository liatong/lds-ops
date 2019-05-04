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

	// Get user value
	//r.GET("/user/:name", handler.GetUser)
	// Insert user 
	r.GET("/insert/:name", handler.InsertUser)
	r.GET("/query/:name", handler.QueryUser)
	r.GET("/delete/:name", handler.DeleteUser)
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	{
		authorized.POST("admin",handler.AuthUser)

	}

	r.POST("/upload",handler.UploadFile)
	r.POST("/package",handler.LishPackage)
	//r.POST("/test/UploadFile",handler.TestUploadFile)


	r.GET("/packagelist",handler.PackageHtml)
	r.GET("/index",handler.IndexHtml)
	r.GET("/code",handler.CodeHtml)
	

	return r
}