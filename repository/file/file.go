package fileRepository

import (
	"file-service/model"
	"github.com/google/uuid"
	"file-service/database"
)

func Create(file model.File) model.File {
	file.ID = uuid.New();
	database.Db.Create(&file);

	return file;
}

func FindAll() []model.File {
	var files []model.File;

	database.Db.Find(&files);

	return files;
}

func FindById(id string) model.File {
	var file model.File;

	database.Db.First(&file, "id = ?", id);

	return file;
}

func DeleteById(id string) model.File {
	var file model.File;

	database.Db.Delete(&file, "id = ?", id);

	return file;
}