package config

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	cfg := LoadConfig()

	dbURL := os.Getenv("DB_URL")
	if dbURL != "" {
		db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err == nil {
			DB = db
			fmt.Println("Connected using DB_URL")
			return db
		}

		log.Println("DATABASE_URL invalid, fallback to manual config:", err.Error())
	}

	dsnRoot := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
	)

	sqlDB, err := sql.Open("mysql", dsnRoot)
	if err != nil {
		panic("Failed to connect to MySQL server.")
	}
	defer sqlDB.Close()

	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = ?)`
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
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database.")
	}

	DB = database
	fmt.Println("Connected using manual config:", cfg.Name)
	return database
}
