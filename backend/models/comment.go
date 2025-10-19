package models

type Comment struct {
	ID           uint `gorm:"primaryKey"`
	CommenterID  uint
	Commenter    User `gorm:"foreignKey:CommenterID"`
	QuoteID      uint
	Quote        Quote `gorm:"foreignKey:QuoteID"`
	LikeCount    uint  `gorm:"default:0" json:"like_count"`
	DislikeCount uint  `gorm:"default:0" json:"dislike_count"`
}
