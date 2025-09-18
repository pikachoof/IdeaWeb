package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type QuoteLike struct {
	ID      uint `gorm:"primaryKey"`
	LikerID uint
	Liker   User `gorm:"foreignKey:LikerID"`
	QuoteID uint
	Quote   Quote `gorm:"foreignKey:QuoteID"`
}

type UpdateQuoteLikeRequest struct {
	LikerID uint `json:"liker_id"`
	QuoteID uint `json:"quote_id"`
}

func AddQuoteLike(db *gorm.DB, likerID uint, quoteID uint) error {
	ctx := context.Background()
	_, err := gorm.G[User](db).Where("id = ?").First(ctx)
	if err != nil {
		return fmt.Errorf("no user found with id %d", likerID)
	}

	quote, err := gorm.G[Quote](db).Where("id = ?").First(ctx)
	if err != nil {
		return fmt.Errorf("no quote found with id %d", quoteID)
	}
	_, err = gorm.G[Quote](db).Where("id = ?", quoteID).Update(ctx, "like_count", quote.LikeCount+1)
	if err != nil {
		return fmt.Errorf("failed to update like count for quote id %d: %v", quoteID, err)
	}
	quoteLike := QuoteLike{
		LikerID: likerID,
		QuoteID: quoteID,
	}
	err = gorm.G[QuoteLike](db).Create(ctx, &quoteLike)
	if err != nil {
		return fmt.Errorf("failed to add like for quote id %d with liker id %d", quoteID, likerID)
	}
	return nil
}

func RemoveQuoteLike(db *gorm.DB, likerID uint, quoteID uint) error {
	ctx := context.Background()
	_, err := gorm.G[User](db).Where("id = ?").First(ctx)
	if err != nil {
		return fmt.Errorf("no user found with id %d", likerID)
	}

	quote, err := gorm.G[Quote](db).Where("id = ?").First(ctx)
	if err != nil {
		return fmt.Errorf("no quote found with id %d", quoteID)
	}
	_, err = gorm.G[Quote](db).Where("id = ?", quoteID).Update(ctx, "like_count", quote.LikeCount-1)
	if err != nil {
		return fmt.Errorf("failed to update like count for quote id %d: %v", quoteID, err)
	}
	_, err = gorm.G[QuoteLike](db).Where("liker_id = ? AND quote_id = ?", likerID, quoteID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to remove like for quote id %d with liker id %d", quoteID, likerID)
	}
	return nil
}
