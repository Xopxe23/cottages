package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %s", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %s", err)
	}
	return &Storage{db}, nil
}

func (s *Storage) PrepareDatabase() error {
	statement, err := s.DB.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36) PRIMARY KEY,
		email VARCHAR(255) UNIQUE,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		password VARCHAR(255),
		is_active INTEGER DEFAULT 1,
		is_verified INTEGER DEFAULT 0,
		is_staff INTEGER DEFAULT 0
	)`)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}
