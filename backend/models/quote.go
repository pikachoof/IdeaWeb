package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

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

func CreateQuote(db *gorm.DB, quote Quote) error {
	ctx := context.Background()
	_, err := gorm.G[Quote](db).Where("id = ?", quote.ID).First(ctx)
	if err == nil {
		return fmt.Errorf("quote with id %d already exists", quote.ID)
	}
	result := gorm.G[Quote](db).Create(ctx, &quote)
	return result
}

func DeleteQuote(db *gorm.DB, quoteID uint) error {
	ctx := context.Background()
	rowsAffected, err := gorm.G[Quote](db).Where("id = ?", quoteID).Delete(ctx)
	if rowsAffected == 0 {
		return fmt.Errorf("no quote found with id %d", quoteID)
	}
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllQuotes(db *gorm.DB) error {
	ctx := context.Background()
	rowsAffected, err := gorm.G[Quote](db).Where("1 = 1").Delete(ctx)
	if rowsAffected == 0 {
		return fmt.Errorf("no quotes to delete")
	}
	if err != nil {
		return err
	}
	return nil
}

func GetQuotes(db *gorm.DB) ([]Quote, error) {
	ctx := context.Background()
	quotes, _ := gorm.G[Quote](db).Find(ctx)
	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quotes found")
	}
	return quotes, nil
}
