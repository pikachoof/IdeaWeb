package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

func CreateUser(db *gorm.DB, user User) error {
	ctx := context.Background()
	_, err := gorm.G[User](db).Where("email = ?", user.Email).First(ctx)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	result := gorm.G[User](db).Create(ctx, &user)
	return result
}

func DeleteUser(db *gorm.DB, userID uint) error {
	ctx := context.Background()
	rowsAffected, err := gorm.G[User](db).Where("id = ?", userID).Delete(ctx)
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", userID)
	}
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllUsers(db *gorm.DB) error {
	ctx := context.Background()
	rowsAffected, err := gorm.G[User](db).Where("1 = 1").Delete(ctx)
	if rowsAffected == 0 {
		return fmt.Errorf("no users to delete")
	}
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(db *gorm.DB) (User, error) {
	ctx := context.Background()
	user, err := gorm.G[User](db).Where("id = ?", user.ID).First(ctx)
	if err == nil {
		return nil, fmt.Errorf("not user found with id %d", user.ID)
	}
	return user, nil
}

func GetUsers(db *gorm.DB) ([]User, error) {
	ctx := context.Background()
	users, _ := gorm.G[User](db).Find(ctx)
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}
	return users, nil
}
