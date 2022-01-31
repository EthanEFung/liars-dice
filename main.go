package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load error", err)
	}

	hub := Hub{
		Lobby:   make(map[*websocket.Conn]bool),
		Rooms:   make(map[string]*Room),
		Channel: make(chan Message),
	}

	http.Handle("/", &Controller{
		Hub: hub,
	})
	go hub.Broadcast()

	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	log.Println("Running Server at ", host+":"+port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
