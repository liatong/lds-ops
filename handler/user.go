package handler

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liatong/lds-ops/model"
	"github.com/liatong/lds-ops/sql"


)

func QueryUser(c *gin.Context){
	/*
	test middler 
 	exampleKey := c.MustGet("example")
 	fmt.Println("Test middle example key:%s",exampleKey)
	*/
	name := c.Params.ByName("name")
	user := model.User{}
	users,err := user.QueryUser(name)
	if err != nil {
		return 
	}
	c.JSON(http.StatusOK,gin.H{"users":users})
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
