package models

type CommentLike struct {
	ID        uint `gorm:"primaryKey"`
	LikerID   uint
	Liker     User `gorm:"foreignKey:LikerID"`
	CommentID uint
	Comment   Comment `gorm:"foreignKey:CommentID"`
}
