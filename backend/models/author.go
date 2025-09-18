package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func AddAuthor(db *gorm.DB, firstName string, lastName string) error {
	ctx := context.Background()
	author := Author{
		FirstName: firstName,
		LastName:  lastName,
	}
	_, err := gorm.G[Category](db).Where("first_name = ? AND last_name", firstName, lastName).First(ctx)
	if err != nil {
		return fmt.Errorf("author with firstname %s and lastname %s already exists")
	}
	err = gorm.G[Author](db).Create(ctx, &author)
	if err != nil {
		return fmt.Errorf("failed to add author with firstname %s and lastname %s", firstName, lastName)
	}
	return nil
}

func RemoveAuthorById(db *gorm.DB, id uint) error {
	ctx := context.Background()
	_, err := gorm.G[Category](db).Where("id = ?", id).First(ctx)
	if err == nil {
		return fmt.Errorf("no author with id %d", id)
	}
	_, err = gorm.G[Author](db).Where("id = ?").Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to remove author with id %d", id)
	}
	return nil
}
