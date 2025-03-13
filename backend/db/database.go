package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
    var err error
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL environment variable not set")
    }

    DB, err = sql.Open("pgx", dbURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if DB == nil {
        log.Fatal("DB is nil after sql.Open")
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Database ping failed:", err)
    }

    log.Println("Connected to PostgreSQL successfully!")
}
