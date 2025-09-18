package models

type Quote struct {
	ID         uint `gorm:"primaryKey"`
	AuthorID   uint
	Author     Author `gorm:"foreignKey:AuthorID"`
	UploaderID uint
	Uploader   User       `gorm:"foreignKey:UploaderID"`
	Text       string     `json:"text"`
	Categories []Category `gorm:"many2many:quote_categories;"`
	LikeCount  uint       `gorm:"default:0" json:"like_count"`
}
