package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/mattes/migrate/migrate"
)

func RunMigrations() {
	migrationsDir := os.Getenv("ET_MIGRATIONS")
	if migrationsDir == "" {
		migrationsDir = "./migrations"
	}

	allErrors, ok := migrate.UpSync(migrateUrl, migrationsDir)
	if !ok {
		for _, err := range allErrors {
			fmt.Println(err)
		}
	}
}

func InitDB() {
	var err error
	db_host := os.Getenv("ET_DB_HOST")
	db_user := os.Getenv("ET_DB_USER")
	db_pass := os.Getenv("ET_DB_PASS")
	db_name := os.Getenv("ET_DB_NAME")

	connection := fmt.Sprintf("host=%v sslmode=disable user=%v password=%v dbname=%v", db_host, db_user, db_pass, db_name)
	migrateUrl = fmt.Sprintf("postgres://%v@%v/%v?sslmode=disable&password=%v", db_user, db_host, db_name, db_pass)

	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to db: %v", db_host)

	db.SetMaxIdleConns(100)
}
