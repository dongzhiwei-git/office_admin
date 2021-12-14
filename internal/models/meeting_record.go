package models

import "time"

type MeetingRecord struct {
	MeetingTime    time.Time `gorm:"meeting_time" json:"meeting_time"`
	JoinPeople     string    `gorm:"join_people" json:"join_people"`
	Record        string    `gorm:"record" json:"record"`
	MeetingContent string    `gorm:"meeting_content" json:"meeting_content"`
}

func (MeetingRecord) TableName() string {
	return "meeting_record"
}
