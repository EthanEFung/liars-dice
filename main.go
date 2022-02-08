package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load error", err)
	}

	hub := Hub{
		Rooms:   make(map[string]*Room),
		lobby:   make(map[*websocket.Conn]*Client),
		channel: make(chan Message),
		rdb: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: "",
			DB:       0,
		}),
	}

	/* redis flush here for development purposes */
	hub.rdb.FlushAll(context.Background())

	http.Handle("/", &Controller{
		Hub: hub,
	})

	go hub.Broadcast()

	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	log.Println("Running Server at ", host+":"+port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
