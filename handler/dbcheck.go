package handler

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"os"
	"time"
	"strings"
	"github.com/liatong/lds-ops/model"


)
const db_uploadBase = "/data/ldsops/dbscript/upload"

func UploadScript(c *gin.Context){

	scriptname := strings.Replace(c.PostForm("scriptname"), " ", "", -1)  
	dbversion := strings.Replace(c.PostForm("dbversion"), " ", "", -1)  
	application := strings.Replace(c.PostForm("appname"), " ", "", -1) 
	version := c.PostForm("appversion")
	branch := c.PostForm("branch")
	comment := c.PostForm("comment")
	file,err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	filename := filepath.Base(file.Filename)
	fileSavePath := db_uploadBase+"/"+application+"/"+version+"/"
	
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

	dbscript := model.Dbscript{scriptname,dbversion,application,version,branch,filefullpath,t,comment}
	err = dbscript.UploadDbcheck()
	if err != nil {
		fmt.Println("Upload dbscript  error:%s",err)
	}
	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully with dbscript=%s  dbversion=%s and app=%s app version=%s.",filename,dbversion, application, version,))
	
}

