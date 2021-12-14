package models

type AttendanceInfo struct {
	ID         string    `gorm:"id" json:"id"`
	Name     string    `gorm:"name" json:"name"`
	Age        int       `gorm:"age" json:"age"`
	Position  string    `gorm:"position" json:"position"`
	Date       string    `gorm:"date" json:"date"`
	Attendance string    `gorm:"attendance" json:"attendance"`
	Recorder   string    `gorm:"recorder" json:"recorder"`
}

func (AttendanceInfo) TableName() string {
	return "attendance_info"
}
