package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "84HW/models"
    _ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
    var err error
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Database connection established")
}

func SaveMessage(msg *models.Message) {
    query := `INSERT INTO messages (email, username, message, timestamp) VALUES ($1, $2, $3, $4)`
    _, err := db.Exec(query, msg.Email, msg.Username, msg.Message, msg.Timestamp)
    if err != nil {
        log.Printf("Error saving message: %v", err)
    }
}
