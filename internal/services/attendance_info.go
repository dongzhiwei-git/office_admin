package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type Attendance struct {
}

func (a *Attendance) GetAttendanceInfo() (store *[]models.AttendanceInfo, err error) {
	AttendanceInfo := new([]models.AttendanceInfo)
	err = dao.Orm.Find(AttendanceInfo).Error

	return AttendanceInfo, err
}
