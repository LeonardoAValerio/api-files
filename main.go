package main

import (
	"github.com/gin-gonic/gin"
	"file-service/controllers"
)

func main() {
	r := gin.Default()
	controllers.InitRoutes(r);
	r.Run(":8080")
}