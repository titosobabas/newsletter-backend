package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"newsletter-backend/types"
	"os"
)

func Connect() (*sql.DB, error) {
	settings := types.DBSettings{
		DatabaseUsername: os.Getenv("DATABASE_USERNAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseNet:      os.Getenv("DATABASE_NET"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
	}

	cfg := mysql.Config{
		User:                 settings.DatabaseUsername,
		Passwd:               settings.DatabasePassword,
		Net:                  settings.DatabaseNet,
		Addr:                 settings.DatabaseHost,
		DBName:               settings.DatabaseName,
		AllowNativePasswords: true,
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(":( NOT Connected!")
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return db, pingErr
}
