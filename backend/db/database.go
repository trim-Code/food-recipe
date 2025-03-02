package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("pgx", "postgres://admin:adminpassword@postgres:5432/foodrecipe")
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
