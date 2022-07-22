package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	InitRoutes(router)
	err := router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
