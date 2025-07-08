package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB : %v", err)

	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB:%v", err)
	}

	log.Println("Connected to Postgres")
	return db
}
