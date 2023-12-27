package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "mariogarza"
    password = ""
    dbname   = "godeck"
)

var (
    db  *sql.DB
    err error
)

type Database struct {
    db *sql.DB
}

func GetDB() *Database{
    return &Database{db: db}
}

// InitDB initializes the database connection pool.
func InitDB() error {
    connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    // Test the database connection
    err = db.Ping()
    if err != nil {
        log.Println(err)
        return err
    }

    log.Println("Connected to the database")

    return nil
}

// CloseDB closes the database connection.
func (d *Database)Close() {
    if d.db != nil {
        d.db.Close()
        log.Println("Database connection closed")
    }
}
