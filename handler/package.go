package handler

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"	
	
	"github.com/liatong/lds-ops/model"
	"github.com/liatong/lds-ops/sql"
	

)

func LishPackage(c *gin.Context){

	
	page,err := strconv.Atoi(c.PostForm("page"))
	if err != nil {
		page = 1 
	}
	pagesize,err  := strconv.Atoi(c.PostForm("pagesize"))
    if err != nil {
    	pagesize = 10
    }

    // 新的通过构造查询语句的工具类来拼凑语句。
    sq := model.SQLQuery{TableName:"ops_package",SQLField:"application_name,env,filename,version,mdcode,filepath,create_time,upload_time"}
	//
	sq.SetLimit(page,pagesize)
    
    // 通过不定查询条件来组合查询语句
    if application := c.PostForm("application"); application != "" {
    	sq.SetWhere("application_name",application)
    }
    if enviroment := c.PostForm("enviroment"); enviroment != "" {
    	sq.SetWhere("env",enviroment)
    }
    if version := c.PostForm("version"); version != "" {
    	sq.SetWhere("version",version)
    }

    //application := c.Params.ByName("application")
    // 原本支持多查询条件的拼凑sql语句的写法。后抽离出来一个工具类来完成
    /*
    sqlCondition := ""
    if application := c.PostForm("application"); application != ""{
    	sqlCondition = sqlCondition + ("application_name='"+application+"'")
    }
    if enviroment := c.PostForm("enviroment"); enviroment != "" {
    	sqlCondition = sqlCondition + (" enviroment="+enviroment)
    }
    if version := c.PostForm("version"); version != "" {
    	sqlCondition = sqlCondition + (" version="+version)
    }
    if sqlCondition != "" {
    	sqlCondition = "where " + sqlCondition 
    }
   
	start := (page -1 ) * pagesize
	sqlCondition = sqlCondition + " limit "+ strconv.Itoa(start) +","+strconv.Itoa(pagesize)
	sqlField := "select application_name,env,filename,version,mdcode,filepath,create_time,upload_time from ops_package "
	sqlScript := sqlField + sqlCondition 
	fmt.Print(sqlScript)
	*/
	
	sqlscript := sq.GetQuery()
	
	rows,err := sql.Pool.Query(sqlscript)
	defer rows.Close()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return 
	}
	packages := make([]model.Package, 0)
	for rows.Next(){
		var p model.Package
		rows.Scan(&p.Application,&p.Enviroment,&p.Name,&p.Version,&p.Mdcode,&p.FilePath,&p.CreateTime,&p.UploadTime)
		packages = append(packages,p)
	}

	if err = rows.Err(); err != nil {
 		return
	}
	fmt.Print(packages)
	
	
	c.JSON(http.StatusOK, gin.H{
	    "status" :200,
	    "error": nil,
	    "data": packages,
	})
}

/*
func ListPackageQuery(c *gin.Context){
	application,err := c.PostForm('application')
	if err != nil {
		sqlScript := sqlScript + 'where application_name = ?' 
	}
	version,err := c.PostForm('version')
	if err != {

	}

}
*/


