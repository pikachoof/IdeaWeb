package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}

func AddCategory(db *gorm.DB, name string) error {
	ctx := context.Background()
	category := Category{
		Name: name,
	}
	_, err := gorm.G[Category](db).Where("name = ?", name).First(ctx)
	if err != nil {
		return fmt.Errorf("category with name %s already exists", name)
	}
	err = gorm.G[Category](db).Create(ctx, &category)
	if err != nil {
		return fmt.Errorf("failed to add category %s", name)
	}
	return nil
}

func RemoveCategoryById(db *gorm.DB, id uint) error {
	ctx := context.Background()
	_, err := gorm.G[Category](db).Where("id = ?", id).First(ctx)
	if err == nil {
		return fmt.Errorf("no category with id %d", id)
	}
	_, err = gorm.G[Category](db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to remove category with id %d", id)
	}
	return nil
}
