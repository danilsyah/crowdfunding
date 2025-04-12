package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password= dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Use db here
	// Get generic database object sql.DB to check the connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance")
	}

	// Ping the database to verify connection
	err = sqlDB.Ping()
	if err != nil {
		panic("failed to ping database")
	}

	fmt.Println("Successfully connected to the database!")
}
