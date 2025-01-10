package lib

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dbPath string) error {
	i, _ := os.Stat(dbPath)
	if i != nil {
		os.Remove(dbPath)
		log.Printf("[Database] Deletion Database - File: %s", dbPath)
	}

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

func AddStaticData(dbPath string, staticDataDir string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	var files []string
	types := []string{
		"agency",
		// "agency_jp",
		"calendar",
		"calendar_dates",
		// "fare_attributes",
		// "fare_rules",
		"feed_info",
		"routes",
		"routes_jp",
		// "shapes",
		"stop_times",
		"stops",
		"trips",
	}
	for _, t := range types {
		files = append(files, staticDataDir+"/"+t+".txt")
	}

	for _, fn := range files {
		f, err := os.Open(fn)
		r := csv.NewReader(f)

		rows, err := r.ReadAll()

		if err != nil {
			db.Close()
			log.Fatal(err)
		}

		var table string
		var column string

		table = strings.TrimSuffix(filepath.Base(fn), filepath.Ext(fn))

		for n, v := range rows {
			var value string

			if n == 0 {
				column = strings.Join(v, ", ")
				continue
			} else {
				for i, val := range v {
					v[i] = fmt.Sprintf("'%s'", val)
				}
				value = strings.Join(v, ", ")
			}

			sentence := fmt.Sprintf("insert into %s (%s) values (%s)", table, column, value)
			log.Printf("[Database] Insert Items - Query: %s", sentence)

			_, err := db.Exec(sentence)
			if err != nil {
				db.Close()
				return err
			}
		}

		defer db.Close()
	}
	return nil
}
