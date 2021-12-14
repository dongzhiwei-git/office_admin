package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type File struct {
}

func (a *File) GetFileInfo() (store *[]models.FileInfo, err error) {
	FileInfo := new([]models.FileInfo)
	err = dao.Orm.Find(FileInfo).Error

	return FileInfo, err
}
