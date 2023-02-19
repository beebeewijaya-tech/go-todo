package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Store interface {
}

type DB struct {
	HOST     string
	USERNAME string
	PASSWORD string
	PORT     int
	DATABASE string
	DIALECT  string
}

func (d *DB) Connect() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.HOST, d.PORT, d.USERNAME, d.PASSWORD, d.DATABASE,
	)
	db, err := sqlx.Connect(d.DIALECT, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(95)

	return db, nil
}
