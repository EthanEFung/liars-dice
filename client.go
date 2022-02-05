package main

import "github.com/gorilla/websocket"

/*
  Client represents servicees
*/
type Client struct {
	Conn *websocket.Conn `json:"-"`
	Name string          `json:"name"`
	Host bool            `json:"host"`
	Room Room            `json:"room"`
	// potentially an IP?
}
