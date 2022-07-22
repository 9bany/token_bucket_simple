package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/hello", middleware, sayHello)
}

func sayHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"messgae": "Hello",
	})
}
