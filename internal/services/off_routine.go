package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Routine struct {
}

func (r *Routine) GetRoutineInfo() (store *[]models.OffRoutine, err error) {
	OffInfo := new([]models.OffRoutine)
	err = dao.Orm.Find(OffInfo).Error
	return OffInfo, err
}
