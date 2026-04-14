package database

import (
	"fmt"
	"log"
	"os"
	"github.com/bakibillahrahat/auth-system/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Create a global variable to hold the database connection.
var DB *gorm.DB

func ConnectDB() {
	// 1. Environment variables for database connection parameters.
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Security Check: Ensure that all required environment variables are set. if any of these variables are missing, log a fatal error and exit the application.
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("❌ Error Database environment variables are missing! Please check your .env or docker-compose file.")
	}

	// 2. Dynamically construct the DSN (Data Source Name) string using the environment variables. This string is used to connect to the PostgreSQL database. The format of the DSN is: "host=HOST" "user=USER" "password=PASSWORD" "dbname=DBNAME" "port=PORT"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	// 3. Use GORM to open a connection to the database using the constructed DSN. If the connection fails, log a fatal error and exit the application.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL Database!")

	// 4. Automatically migrate the User model to create the users table in the database if it doesn't already exist. This ensures that the database schema is up to date with the User model defined in the internal/models package.
	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("❌ Failed to migrate User model:", err)
	}
	fmt.Println("✅ Successfully migrated User model!")

	// 5. Assign the database connection to the global variable DB so that it can be accessed throughout the application.
	DB = db
}