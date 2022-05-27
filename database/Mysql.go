package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DATABASE_USERNAME"),
		Passwd:               os.Getenv("DATABASE_PASSWORD"),
		Net:                  os.Getenv("DATABASE_NET"),
		Addr:                 os.Getenv("DATABASE_HOST"),
		DBName:               os.Getenv("DATABASE_NAME"),
		AllowNativePasswords: true,
	}
	fmt.Println("os.Getenv(\"DATABASE_USERNAME\") ----> ", os.Getenv("DATABASE_USERNAME"))
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return db, pingErr
}
