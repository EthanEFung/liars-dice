package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Controller struct {
	Hub Hub
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Could not upgrade connection: ", err)
	}
	c.Hub.lobby[conn] = &Client{}
	c.Hub.Publish(conn)
}
