package services

import (
	"inherited/internal/dao"
	"inherited/internal/models"
)

type MeetingRecord struct {
}

func (a *MeetingRecord) GetFileInfo() (store *[]models.MeetingRecord, err error) {
	MeetingRecordInfo := new([]models.MeetingRecord)
	err = dao.Orm.Find(MeetingRecordInfo).Error
	return MeetingRecordInfo, err
}
