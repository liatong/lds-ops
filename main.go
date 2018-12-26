package main

import (
/*
	"net/http"


	"github.com/gin-gonic/gin"
*/
	"github.com/liatong/lds-ops/router"

)

func main() {
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9091")
}
