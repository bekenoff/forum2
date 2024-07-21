package storage

import (
	"database/sql"
	"forum/pkg/config"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqlite(config config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DB.Driver, config.DB.Dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = CreateTables(db, config); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB, config config.Config) error {
	file, err := os.ReadFile(config.DB.Database)
	if err != nil {
		return err
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		request = strings.TrimSpace(request)
		if request == "" {
			continue
		}
		log.Printf("Executing SQL statement: %s", request)
		_, err := db.Exec(request)
		if err != nil {
			log.Printf("Error executing SQL statement: %s, error: %v", request, err)
			return err
		}
	}
	return nil
}
