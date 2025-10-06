package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Comment struct {
	ID          uint `gorm:"primaryKey"`
	CommenterID uint
	Commenter   User `gorm:"foreignKey:CommenterID"`
	QuoteID     uint
	Quote       Quote `gorm:"foreignKey:QuoteID"`
	LikeCount   uint  `gorm:"default:0" json:"like_count"`
}

func CreateComment(db *gorm.DB, comment Comment) error {
	ctx := context.Background()
	_, err := gorm.G[Comment](db).Where("id = ?", comment.ID).First(ctx)
	if err == nil {
		return fmt.Errorf("comment with id %d already exists", comment.ID)
	}
	result := gorm.G[Comment](db).Create(ctx, &comment)
	return result
}

func DeleteComment(db *gorm.DB, commentID uint) error {
	ctx := context.Background()
	rowsAffected, err := gorm.G[Comment](db).Where("id = ?", commentID).Delete(ctx)
	if rowsAffected == 0 {
		return fmt.Errorf("no comment found with id %d", commentID)
	}
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllComments(db *gorm.DB) error {
	ctx := context.Background()
	rowsAffected, err := gorm.G[Comment](db).Where("1 = 1").Delete(ctx)
	if rowsAffected == 0 {
		return fmt.Errorf("no comments to delete")
	}
	if err != nil {
		return err
	}
	return nil
}

func GetComments(db *gorm.DB) ([]Comment, error) {
	ctx := context.Background()
	comments, _ := gorm.G[Comment](db).Find(ctx)
	if len(comments) == 0 {
		return nil, fmt.Errorf("no comments found")
	}
	return comments, nil
}
