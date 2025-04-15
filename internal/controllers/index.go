package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoavalerio/api-files/internal/controllers/files"
)

func InitRoutes(r *gin.Engine) {
	filesController.InitRoutes(r);
}