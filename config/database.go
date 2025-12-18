package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping db:", err)
	}

	log.Println("✅ Database Connected")
	return db
}
