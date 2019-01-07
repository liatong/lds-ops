package handler

import (
	"net/http"
	"fmt"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
)



func Pong(c *gin.Context) {
		c.String(http.StatusOK, "Hello word! Ping Pong!")
}

func IndexHtml(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html","")
}
func CodeHtml(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html","")
}

var db = make(map[string]string)

func GetUser(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
}

func AuthUser(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Print(user)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
}
func IsDir(name string) bool {
    fi, err := os.Stat(name)
    if err != nil {
        fmt.Println("Error: ", err)
        return false
    }

    return fi.IsDir()
}

func UploadFile(c *gin.Context){
	//application := c.PostForm("application")
	//version := c.PostForm("version")
	application := "admin-service"
	version := "1.0"

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	filepath := "/tmp/upload/"+application+"/"+version

	if !IsDir(filepath){
		err = os.MkdirAll(filepath,0755)
		if err != nil {
	    	c.String(http.StatusBadRequest, fmt.Sprintf("can't mkdir filepath:%s err: %s",filepath,err.Error()))
			return
	    }
	}
	
	filefullpath := filepath+"/"+filename

	fmt.Print(filefullpath)
	if err := c.SaveUploadedFile(file, filefullpath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields application=%s and version=%s.",filename, application, version))
}




