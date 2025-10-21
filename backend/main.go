package main

import (
	"IdeaWeb/initializations"
	"IdeaWeb/models"
	"fmt"
	"log"
	"os"
	"strconv"

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
	router.Run()
}
