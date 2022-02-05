package main

/*
  TODOS:
	- [ ] figure out how to show the client some sort of profile (could just be username for now) of
		people who are currently occupying the room
*/

import (
	"context"

	"github.com/gorilla/websocket"
)

type Room struct {
	Name     string           `json:"name"`
	Hostname string           `json:"hostname"`
	Clients  map[*Client]bool `json:"-"`
	Hub      *Hub             `json:"-"`
	Channel  chan Message     `json:"-"`
}

func (r Room) Join(ctx context.Context) {

}

func (r Room) Publish(ctx context.Context, conn *websocket.Conn) {}

func (r Room) Subscribe(ctx context.Context, conn *websocket.Conn) {}
