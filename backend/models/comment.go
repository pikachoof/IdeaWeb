package models

type Comment struct {
	ID          uint `gorm:"primaryKey"`
	CommenterID uint
	Commenter   User `gorm:"foreignKey:CommenterID"`
	QuoteID     uint
	Quote       Quote `gorm:"foreignKey:QuoteID"`
}
