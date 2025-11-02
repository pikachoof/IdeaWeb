package services

import (
	"IdeaWeb/models"
)

type UserServiceInterface interface {
	Create(newUser *models.User) error
	GetAll() ([]*models.User, error)
	GetByID(id uint) (*models.User, error)
	SetRole(id uint, newRole *models.UserRole)
	Update(id uint, updatedUser *models.User) error
	DeleteByID(id uint) error
}
