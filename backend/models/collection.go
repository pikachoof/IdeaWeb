package models

type Collection struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"name"`
	UserID uint   `gorm:"foreignKey:UserID"`
	User   User   `gorm:"foreignKey:UserID"`
}
