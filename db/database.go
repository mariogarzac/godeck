package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

func (d *Database)RegisterUser(username, email, password string) error {

    // check if user already exists
    stmt := "SELECT email FROM Users WHERE email = $1"
    row := db.QueryRow(stmt, email)

    var isDuplicate string
    err = row.Scan(&isDuplicate)

    if err != sql.ErrNoRows {
        log.Println("Error user already exists", err)
        return err
    }

    // insert user into table
    var insert *sql.Stmt

    insert, err := db.Prepare("INSERT into Users (username, email, password) VALUES ($1, $2, $3)")
    if err != nil {
        log.Println("Error preparing query", err)
        return err
    }
    defer insert.Close()

    var hashedPassword []byte
    hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error creating the hash ", err)
        return err
    }

    log.Println(len(hashedPassword))

    _, err = insert.Exec(username, email, hashedPassword)

    if err != nil {
        log.Println("Error inserting data", err)
        return err
    }

    return nil
}

