package main

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load() // Load ENV Variables

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitDB()    // Start DB
	MigrateDB() // Run migrations
}

func main() {
	Run()
}
