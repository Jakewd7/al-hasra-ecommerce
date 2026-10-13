package config

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	cfg := LoadConfig()

	dbURL := os.Getenv("DB_URL")
	if dbURL != "" {
		db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err == nil {
			DB = db
			fmt.Println("Connected using DB_URL")
			return db
		}

		log.Println("DATABASE_URL invalid, fallback to manual config:", err.Error())
	}

	dsnRoot := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
	)

	sqlDB, err := sql.Open("postgres", dsnRoot)
	if err != nil {
		panic("Failed to connect to PostgreSQL server.")
	}
	defer sqlDB.Close()

	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)`
	err = sqlDB.QueryRow(query, cfg.Name).Scan(&exists)
	if err != nil {
		panic("Error checking database: " + err.Error())
	}

	if !exists {
		fmt.Printf("Database '%s' does not exist. Create it? (y/n): ", cfg.Name)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "y" || input == "yes" {
			_, err := sqlDB.Exec("CREATE DATABASE " + cfg.Name)
			if err != nil {
				panic("Failed to create database: " + err.Error())
			}
			fmt.Println("Database created successfully.")
		} else {
			panic("Database not found. Operation cancelled.")
		}
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Pass,
		cfg.Name,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database.")
	}

	DB = database
	fmt.Println("Connected using manual config:", cfg.Name)
	return database
}
