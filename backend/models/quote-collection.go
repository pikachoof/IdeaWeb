package models

type CollectionQuote struct {
	ID           uint `gorm:"primaryKey"`
	CollectionID uint
	Collection   Collection `gorm:"foreignKey:CollectionID"`
	QuoteID      uint
	Quote        Quote `gorm:"foreignKey:QuoteID"`
}
