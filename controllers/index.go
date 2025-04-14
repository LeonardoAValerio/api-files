package controllers

import (
	"github.com/gin-gonic/gin"
	"file-service/controllers/files"
)

func InitRoutes(r *gin.Engine) {
	filesController.InitRoutes(r);
}