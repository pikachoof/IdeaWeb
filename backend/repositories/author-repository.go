package repositories

import (
	"IdeaWeb/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func (r *AuthorRepository) GetAuthorByID(authorID uint) (models.Author, error) {
	ctx := context.Background()
	author, err := gorm.G[models.Author](r.db).Where("id = ?", authorID).First(ctx)
	if err != nil {
		return models.Author{}, fmt.Errorf("author with id = %d doesn't exist", authorID)
	}
	return author, nil
}

func (r *AuthorRepository) GetAllAuthors() ([]models.Author, error) {
	ctx := context.Background()
	authors, err := gorm.G[models.Author](r.db).Where("1 = 1")
	if err != nil {
		return []models.Author{}, fmt.Error
	}
}
