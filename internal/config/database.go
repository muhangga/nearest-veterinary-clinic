package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func OpenDB(dsn string, setLimits bool) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if setLimits {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	fmt.Println("Connected to database")
	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func RollbackDB(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		log.Fatal(err)
	}
}
