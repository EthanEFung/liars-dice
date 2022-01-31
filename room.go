package main

import (
	"context"

	"github.com/gorilla/websocket"
)

type Room struct {
	Name     string                   `json:"name"`
	Hostname string                   `json:"hostname"`
	Clients  map[*websocket.Conn]bool `json:"-"`
	Hub      Hub                      `json:"-"`
	Channel  chan Message             `json:"-"`
}

func (r Room) Join(ctx context.Context, u User) {}

func (r Room) Publish(ctx context.Context, conn *websocket.Conn) {}

func (r Room) Subscribe(ctx context.Context, conn *websocket.Conn) {}
