package lib

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dbPath string) error {
	log.Printf("[Database] Initialize Database - File: %s", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlBytes, err := os.ReadFile("./lib/static.sql")
	if err != nil {
		return err
	}
	sqlString := string(sqlBytes)

	_, err = db.Exec(sqlString)
	if err != nil {
		return err
	}
	return nil
}
