package main

import (
	"9bany/rate-limiter-token-bucket/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.InitRoutes(router)
	err := router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
