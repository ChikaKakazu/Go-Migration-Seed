package main

import (
	"ChikaKakazu/Go-Migration-Seed/database/seeds/seed"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := OpenDB(os.Getenv("DRIVER"), os.Getenv("DSN"))
	defer CloseDB(db)

	if err := db.Ping(); err != nil {
		log.Fatal("db.Ping failed : ", err)
	}

	// シードする
	for _, seed := range seed.All() {
		fmt.Println("table name : ", seed.Name)
		seed.Run(db)
	}
}

func OpenDB(driver, dsn string) *sql.DB {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("OpenDB failed:", err)
	}
	return db
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal("CloseDB failed:", err)
	}
}
