package handler

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"os"
	"time"
	"strings"
	"strconv"
	"github.com/liatong/lds-ops/model"
	"github.com/liatong/lds-ops/sql"
	


)
const db_uploadBase = "/data/ldsops/dbscript/upload/"

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
	SavePath := db_uploadBase+"/"+application+"/"+version+"/"
	
	if !IsDir(SavePath){
		//fmt.Println("SavePath no exit, need create!")
		err = os.MkdirAll(SavePath,0755)
		if err != nil {
	    	c.String(http.StatusBadRequest, fmt.Sprintf("can't mkdir filepath:%s err: %s",SavePath,err.Error()))
			return
	    }
	    //fmt.Println("SavePath:%s create successfully!",SavePath)
	}
	filefullpath := SavePath+filename
	if err := c.SaveUploadedFile(file, filefullpath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
	}
	dbfilepath := application+"/"+version+"/"+filename
	
	t := time.Now().Format("2006-01-02 15:04:05.000")
    
	dbscript := model.Dbscript{scriptname,dbversion,application,version,branch,dbfilepath,t,comment}
	err = dbscript.UploadDbcheck()
	if err != nil {
		fmt.Println("Upload dbscript  error:%s",err)
	}
	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully with dbscript=%s  dbversion=%s and app=%s app version=%s.",filename,dbversion, application, version,))
	
}

func LishDbscript(c *gin.Context){

	page,err := strconv.Atoi(c.PostForm("page"))
	if err != nil {
		page = 1 
	}
	pagesize,err  := strconv.Atoi(c.PostForm("pagesize"))
    if err != nil {
    	pagesize = 10
    }

    // 新的通过构造查询语句的工具类来拼凑语句。
    sq := model.SQLQuery{TableName:"ops_dbscript",SQLField:"scriptname,version,app,appversion,branch,filepath,createtime,comment"}
	//
	sq.SetLimit(page,pagesize)
    
    // 通过不定查询条件来组合查询语句
    if scriptname := c.PostForm("scriptname"); scriptname != "" {
    	sq.SetWhere("scriptname",scriptname)
    }
    if version := c.PostForm("version"); version != "" {
    	sq.SetWhere("version",version)
    }
    if app := c.PostForm("app"); app != "" {
    	sq.SetWhere("app",app)
    }
    if appversion := c.PostForm("appversion"); appversion != "" {
    	sq.SetWhere("appversion",appversion)
    }

	
	sqlscript := sq.GetQuery()
	
	rows,err := sql.Pool.Query(sqlscript)
	defer rows.Close()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return 
	}
	dbscripts := make([]model.Dbscript, 0)
	for rows.Next(){
		var p model.Dbscript
		//rows.Scan(&p.Application,&p.Enviroment,&p.Name,&p.Version,&p.Mdcode,&p.FilePath,&p.CreateTime,&p.UploadTime)
		rows.Scan(&p.Name,&p.DbVersion,&p.App,&p.AppVersion,&p.Branch,&p.FilePath,&p.CreateTime,&p.Comment)
		dbscripts = append(dbscripts,p)
	}

	if err = rows.Err(); err != nil {
 		return
	}
	fmt.Print(dbscripts)
	
	
	c.JSON(http.StatusOK, gin.H{
	    "status" :200,
	    "error": nil,
	    "data": dbscripts,
	})
}

