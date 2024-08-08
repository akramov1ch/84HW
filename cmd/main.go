package main

import (
	"84HW/db"
	"84HW/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/ws", handlers.HandleWebSocket)

	log.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
