package main

import (
	"IdeaWeb/db"
	"IdeaWeb/handlers"
	"IdeaWeb/middleware"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	// load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	loadEnv()

	log.Print("Hello Logger!")

	db, err := db.ConnectDB(
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
	router.Use(middleware.AuthMiddleware)

	admin := router.Group("/admin")
	admin.Use(middleware.AdminRoleMiddleware)
	{
		adminUsers := admin.Group("/users")
		{
			adminUsers.GET("", handlers.GetAllUsers)
			adminUsers.GET("/:id", handlers.GetUser)
			adminUsers.PATCH("/:id/set-admin", handlers.SetAdminUser)
			adminUsers.PATCH("/:id/set-regular-user", handlers.SetRegularUser)
		}

		adminQuotes := admin.Group("/quotes")
		{
			adminQuotes.GET("", handlers.GetAllQuotes)
			adminQuotes.GET("/:id", handlers.GetQuote)
			adminQuotes.PATCH("/:id/approve", handlers.ApproveQuote)
			adminQuotes.PATCH("/:id/reject", handlers.RejectQuote)
		}
	}

	user := router.Group("/user")
	user.Use(middleware.UserRoleMiddleware)
	{
		userQuotes := user.Group("/quotes")
		{
			userQuotes.GET("", handlers.GetAllUserQuotes)
			userQuotes.GET("/:id", handlers.GetUserQuote)
			userQuotes.POST("/submissions", handlers.SubmitQuote)
			userQuotes.DELETE("/submissions/:id", handlers.RemoveSubmission)
			userQuotes.DELETE("/:id", handlers.DeleteUserQuote)
		}
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
