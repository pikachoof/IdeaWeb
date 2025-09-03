package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
What do I need for this file?
1) Get all quotes
2) Create a quote
3) Delete a quote

That's it.

User: (Name, Surname, Email, Password)
Quote: (Author[Author[]], Uploader[User], Text, Categories[Categories[]])
Author: (Name)
Category: (Name)
*/

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// I need to test changing the names of the foreign keys here. (Will it still preserve the correct relationship?)
type Quote struct {
	ID         int `json:"id"`
	AuthorID   int
	Author     Author `gorm:"foreignKey:AuthorID"`
	UploaderID int
	Uploader   User       `gorm:"foreignKey:UploaderID"`
	Text       string     `json:"text"`
	Categories []Category `gorm:"many2many:quote_categories;"`
}

func loadEnv() {
	// load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func connectDB(host, user, password, dbname, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createUser(db *gorm.DB, user User) error {
	ctx := context.Background()
	_, err := gorm.G[User](db).Where("email = ?", user.Email).First(ctx)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	result := gorm.G[User](db).Create(ctx, &user)
	return result
}

func deleteUser(db *gorm.DB, userID int) error {
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

func deleteAllUsers(db *gorm.DB) error {
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

func getUsers(db *gorm.DB) ([]User, error) {
	ctx := context.Background()
	users, _ := gorm.G[User](db).Find(ctx)
	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}
	return users, nil
}

func main() {
	loadEnv()

	db, err := connectDB(
		os.Getenv("POSTGRES_HOSTNAME"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	if err != nil {
		panic("failed to connect to the database")
	}

	fmt.Println("Successfully connected to the database")

	/*
		migrationError := db.AutoMigrate(&User{}, &Author{}, &Category{}, &Quote{})
		if migrationError != nil {
			panic("failed to migrate database")
		}

		fmt.Println("Database migrated successfully")
	*/

	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		users, err := getUsers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch users: " + err.Error()})
			return
		}
		c.JSON(200, users)
	})

	router.POST("/users/create", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		err := createUser(db, newUser)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user: " + err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User created successfully"})
	})

	router.DELETE("/users/delete/:id", func(c *gin.Context) {
		userID, convErr := strconv.Atoi(c.Param("id"))
		if convErr != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		err := deleteUser(db, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete user: " + err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User deleted successfully"})
	})

	router.DELETE("/users/delete-all", func(c *gin.Context) {
		err := deleteAllUsers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete all users: " + err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "All users deleted successfully"})
	})

	router.Run()
}
