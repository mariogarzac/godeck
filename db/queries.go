package db

import (
    "database/sql"
    "fmt"
    "log"
    "strings"

    _ "github.com/lib/pq"
    "github.com/mariogarzac/godeck/pkg/models"
    "golang.org/x/crypto/bcrypt"
)

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

    _, err = insert.Exec(username, email, hashedPassword)

    if err != nil {
        log.Println("Error inserting data", err)
        return err
    }

    return nil
}

func (d *Database) SaveFileToDB(fileType, fileName string, fileSize int64) error {

    stmt := `INSERT INTO sound_library (user_id, sound_path, file_ext, file_size) 
    VALUES ($1, $2, $3, $4)`

    insert, err := db.Prepare(stmt)

    if err != nil {
        log.Println("Error preparing query", err)
        return err
    }

    defer insert.Close()

    _, err = insert.Exec(1, fileName, fileType, fileSize)

    if err != nil {
        log.Println("Error inserting data", err)
        return err
    }

    return nil
}

func (d *Database) GetFilesByUserId(userId int) ([]models.File, error) {

    var files []models.File

    stmt := "SELECT sound_path, file_ext, file_size FROM sound_library WHERE user_id = $1"

    rows, err := db.Query(stmt, userId)
    fmt.Println("rows are", rows)

    if err != nil {
        log.Println(err)
        return files, err
    }

    var fileName, fileType, fileSize string

    fmt.Println("entre")
    for rows.Next() {
        if err := rows.Scan(&fileName, &fileType, &fileSize); err != nil {
            log.Println(err)
            return files, err
        }

        name := strings.Split(fileName, "/")

        file := models.File{
            FileName: name[len(name) - 1],
            FileType: fileType,
        }

        files = append(files, file)

    }

    return files, nil
}
