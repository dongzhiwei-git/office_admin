package models

type OffRoutine struct {
	ID          int    `json:"id" json:"id"`
	Time        string `gorm:"time" json:"time"`
	Transaction string `gorm:"transaction" json:"transaction"`
	Record      string `gorm:"record" json:"record"`
}

func (OffRoutine) TableName() string {
	return "off_routine"
}
