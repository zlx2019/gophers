package store

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewStore(c mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("DB: Successfully connected!")
	return db, nil
}
