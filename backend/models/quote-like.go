package models

type QuoteLike struct {
	ID      uint `gorm:"primaryKey"`
	LikerID uint
	Liker   User `gorm:"foreignKey:LikerID"`
	QuoteID uint
	Quote   Quote `gorm:"foreignKey:QuoteID"`
	IsLike  bool  `gorm:"not null;default:false"`
}

type UpdateQuoteLikeRequest struct {
	LikerID uint `json:"liker_id"`
	QuoteID uint `json:"quote_id"`
}
