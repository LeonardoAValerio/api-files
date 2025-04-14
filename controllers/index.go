package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoavalerio/api-files/controllers/files"
)

func InitRoutes(r *gin.Engine) {
	filesController.InitRoutes(r);
}