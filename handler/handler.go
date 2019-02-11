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
	"github.com/liatong/lds-ops/sql"

)

/*
func newPool() *sql.DB {
	
	db, err := sql.Open("mysql","root:Password@tcp(127.0.0.1:3306)/lds_ops")
	if err != nil {
		fmt.Println("mysql error")
	}
	if err := db.Ping(); err != nil {
		fmt.Println("mysql error")
	}
	return db
}

var pool = newPool()


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
//---- 测试database/sql ----//
type User struct {
	ID   int64
	Name string
}


func InsertUser(c *gin.Context){
	name := c.Params.ByName("name")
	// 测试插入用户信息
	res, err := sql.Pool.Exec("insert into `test_name` (`name`) values (?)", name)
	fmt.Println(res.LastInsertId()) 
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("can't insert into mysql: %s",err.Error()))
		return
	}
	
	c.String(http.StatusOK, fmt.Sprintf("name:%s",name))
}

func DeleteUser(c *gin.Context){
	name := c.Params.ByName("name")
	_, err := sql.Pool.Exec("delete from `test_name` where `name` = ?", name)
	if err != nil {
		return 
	}
	c.String(http.StatusOK, fmt.Sprintf("delete user:%s",name))
}


//-----------完成database/sql数据库操作测试--------//
func UploadFile(c *gin.Context){
	application := c.PostForm("application")
	version := c.PostForm("version")
	//application := "admin-service"
	//version := "1.0"

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	filepath := uploadBase+"/"+application+"/"+version

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
	t := time.Now()
	uploadTime := t.Unix()
	fileMd5,_ := GetFileMd5(filefullpath)

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields application=%s and version=%s. md5sum=%s, uploadTime=%d",filename, application, version,fileMd5,uploadTime))
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


