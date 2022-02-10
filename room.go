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
	Name     string                      `json:"name"`
	Hostname string                      `json:"hostname"`
	clients  map[*websocket.Conn]*Client `json:"-"`
	hub      *Hub                        `json:"-"`
	channel  chan Message                `json:"-"`
}

func (r *Room) Join(ctx context.Context, msg Message, conn *websocket.Conn) Message {
	var p struct {
		Username string `json:"username"`
	}
	msg.DecodePayload(&p)
	clients := r.clients
	hostless := len(clients) == 0

	if hostless {
		r.Hostname = p.Username
	}
	for wc, client := range r.hub.lobby {
		if wc == conn {
			client.Host = hostless
			client.Room = r
			r.clients[wc] = client
			break
		}
	}
	return Message{
		Type:    JoinedType,
		Payload: r,
	}
}

func (r *Room) Leave(ctx context.Context, msg Message) Message {
	var p struct {
		Username string `json:"username"`
	}
	msg.DecodePayload(&p)
	for wc, client := range r.clients {
		if client.Name == p.Username {
			delete(r.clients, wc)
		}
	}
	return Message{
		Type:    LeftType,
		Payload: r,
	}
}

func (r Room) Publish(ctx context.Context, conn *websocket.Conn) {}

func (r Room) Subscribe(ctx context.Context, conn *websocket.Conn) {}
