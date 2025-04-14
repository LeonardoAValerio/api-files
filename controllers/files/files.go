package filesController

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoavalerio/api-files/repository/file"
	"github.com/leonardoavalerio/api-files/model"
	"path/filepath"
	"fmt"
	"net/http"
	"os"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/files", func(c *gin.Context) {
		c.JSON(200, fileRepository.FindAll())
	});

	r.POST("/files", func(c *gin.Context) {
		file, err := c.FormFile("file");
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo não encontrado"})
			return
		}

		filename := filepath.Base(file.Filename)
		typeFile := filepath.Ext(file.Filename)
		
		createdFile := fileRepository.Create(model.File{Filename: filename, Size: file.Size, Type: typeFile});
		
		savePath := fmt.Sprintf("./uploads/%s.%s", createdFile.ID, createdFile.Type)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Upload feito com sucesso",
			"file":    createdFile,
		});
	});

	r.GET("/files/:id", func(c *gin.Context) {
		id := c.Param("id");

		file := fileRepository.FindById(id);
		filePath := fmt.Sprintf("./uploads/%s.%s", file.ID, file.Type);

		c.File(filePath);
	});

	r.DELETE("/files/:id", func(c *gin.Context) {
		id := c.Param("id");

		file := fileRepository.FindById(id);
		filePath := fmt.Sprintf("./uploads/%s.%s", file.ID, file.Type);

		if err := os.Remove(filePath); err != nil {
			c.JSON(500, gin.H{"error": "Erro ao deletar o arquivo"})
			return
		}

		fileRepository.DeleteById(id);
	
		c.JSON(200, gin.H{"message": "Arquivo deletado com sucesso"})
	});
}