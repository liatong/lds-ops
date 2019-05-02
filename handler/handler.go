package handler

import (
	"net/http"
	"fmt"
	"os"
	"crypto/md5"
	"encoding/hex"
	"io"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/liatong/lds-ops/model"
)

/*

var pool *sql.DB
func init() {
	var err error
	pool, err = sql.Open("mysql","root:Password@tcp(127.0.0.1:3306)/lds_ops")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = pool.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
*/

const uploadBase = "/tmp/upload"

/*  load  html web */
func Pong(c *gin.Context) {
		c.String(http.StatusOK, "Hello word! Ping Pong!")
}

func IndexHtml(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html","")
}
func CodeHtml(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html","")
}


/************ Handler *********/
var db = make(map[string]string)

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

// check file path if directory.
func IsDir(name string) bool {
    fi, err := os.Stat(name)
    if err != nil {
      //This path is not exit or this path is a file! 
      return false
    }

    return fi.IsDir()
}

//-----------完成database/sql数据库操作测试--------//
func UploadFile(c *gin.Context){

	fmt.Println("-->debug upload file!")
	application := c.PostForm("application")
	version := c.PostForm("version")
	env := c.PostForm("enviroment")
	file,err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	
	filename := filepath.Base(file.Filename)
	fileSavePath := uploadBase+"/"+application+"/"+version+"/"
	
	if !IsDir(fileSavePath){
		//fmt.Println("fileSavePath no exit, need create!")
		err = os.MkdirAll(fileSavePath,0755)
		if err != nil {
	    	c.String(http.StatusBadRequest, fmt.Sprintf("can't mkdir filepath:%s err: %s",fileSavePath,err.Error()))
			return
	    }
	    //fmt.Println("fileSavePath:%s create successfully!",fileSavePath)
	}
	filefullpath := fileSavePath+filename
	if err := c.SaveUploadedFile(file, filefullpath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
	}
	
	t := time.Now().Format("2006-01-02 15:04:05.000")
	fileMd5,_ := GetFileMd5(filefullpath)

	/*
	name := "liwentong"
	user := model.User{}
	users,err := user.QueryUser(name)
	if err != nil {
		return 
	}
	fmt.Println(users)
	*/

	pk := model.Package{filename,application,env,version,fileMd5,filefullpath,t,t}
	fmt.Println(pk)
	err = pk.UploadPackage()
	if err != nil {
		fmt.Println("Upload package database error:%s",err)
	}
	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields application=%s and version=%s. md5sum=%s, uploadTime=%d",filename, application, version,fileMd5,t))
	
}


func GetFileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("os Open error")
		return "", err
	}
	defer file.Close()

	md5 := md5.New()
	_, err = io.Copy(md5, file)
	if err != nil {
		fmt.Println("io copy error")
		return "", err
	}
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str, nil
}


