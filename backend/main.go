package main

import (
	"IdeaWeb/initializations"
	"IdeaWeb/models"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func loadEnv() {
	// load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	loadEnv()

	db, err := initializations.ConnectDB(
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
		users, err := models.GetUsers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch users: " + err.Error()})
			return
		}
		c.JSON(200, users)
	})

	router.POST("/users/create", func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		err := models.CreateUser(db, newUser)
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
		err := models.DeleteUser(db, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete user: " + err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User deleted successfully"})
	})

	router.DELETE("/users/delete-all", func(c *gin.Context) {
		err := models.DeleteAllUsers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete all users: " + err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "All users deleted successfully"})
	})

	router.POST("quotes/like/:id", func(c *gin.Context) {
		var req models.UpdateQuoteLikeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		err := models.AddQuoteLike(db, req.LikerID, uint(req.QuoteID))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to like quote: " + err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Quote liked successfully"})
	})

	router.POST("quotes/dislike/:id", func(c *gin.Context) {
		var req models.UpdateQuoteLikeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		err := models.RemoveQuoteLike(db, req.LikerID, uint(req.QuoteID))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to dislike quote: " + err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Quote disliked successfully"})
	})
}
