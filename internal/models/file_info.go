package models

type FileInfo struct {
	FileNumber  string `gorm:"file_number" json:"file_number"`
	FileKind    string `gorm:"file_kind" json:"file_kind"`
	FileName    string `gorm:"file_name" json:"file_name"`
	StoreRecord string `gorm:"store_record" json:"store_record"`
	Record      string `gorm:"record" json:"record"`
}

func (FileInfo) TableName() string {
	return "file_info"
}
