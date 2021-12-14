package models

import "time"

type OffRoutine struct {
	id          int       `json:"id" json:"id"`
	time        time.Time `gorm:"meeting_time" json:"meeting_time"`
	transaction string    `gorm:"join_people" json:"join_people"`
	Record      string    `gorm:"record" json:"record"`
}

func (OffRoutine) TableName() string {
	return "off_routine"
}
