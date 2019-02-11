package handler

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/liatong/lds-ops/model"

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
