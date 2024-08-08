package models

import "time"

type Message struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Username  string    `json:"username"`
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}
