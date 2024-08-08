package handlers

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    "84HW/models"
    "84HW/db"
    "time"
)

var clients = make(map[*websocket.Conn]bool) 
var broadcast = make(chan models.Message)    

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    clients[ws] = true

    for {
        var msg models.Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("error: %v", err)
            delete(clients, ws)
            break
        }
        msg.Timestamp = time.Now()

        db.SaveMessage(&msg)

        broadcast <- msg
    }
}

func HandleMessages() {
    for {
        msg := <-broadcast
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("error: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}
