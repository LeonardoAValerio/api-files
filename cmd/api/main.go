package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoavalerio/api-files/internal/controllers"
)

func main() {
	r := gin.Default()
	controllers.InitRoutes(r);
	r.Run(":8080")
}