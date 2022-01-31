package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Lobby   map[*websocket.Conn]bool `json:"-"`
	Rooms   map[string]*Room         `json:"rooms"`
	Channel chan Message             `json:"-"`
}

/*
	Publish reads the messages coming from the clients websocket connection and sends a Message
	to the Hub Channel
*/
func (h *Hub) Publish(conn *websocket.Conn) {
	for {
		var msg Message
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Publish error: ", err)
			delete(h.Lobby, conn)
			break
		}
		switch msg.Type {
		case ConnectType:
			h.Channel <- h.GetRooms(ctx)
		case CreateType:
			h.Channel <- h.NewRoom(ctx, msg.Payload)
			h.Channel <- h.GetRooms(ctx)
		case JoinType:
			h.Channel <- h.JoinRoom(ctx, msg.Payload)
		case LeaveType:
			h.Channel <- h.LeaveRoom(ctx, msg.Payload)
		}
	}
}

/*
	Broadcast receives Messages from the hub's Channel and writes the message to each
	connection within the Lobby
*/
func (h *Hub) Broadcast() {
	for {
		msg := <-h.Channel
		for conn, ok := range h.Lobby {
			// here we can have some sort of middleware that prevents
			// certain events from firing to specific clients
			if err := conn.WriteJSON(&msg); !ok || (err != nil && !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF) {
				log.Println("Unmarshal error: ", err)
				conn.Close()
				delete(h.Lobby, conn)
				continue
			}
		}
	}
}

func (h *Hub) GetRooms(ctx context.Context) Message {
	var rooms []Room
	return Message{
		Type:    RoomsType,
		Payload: rooms,
	}
}

func (h *Hub) NewRoom(ctx context.Context, payload interface{}) Message {
	var room Room
	return Message{
		Type:    CreatedType,
		Payload: room,
	}
}

func (h *Hub) JoinRoom(ctx context.Context, payload interface{}) Message {
	var room Room
	return Message{
		Type:    JoinedType,
		Payload: room,
	}
}

func (h *Hub) LeaveRoom(ctx context.Context, payload interface{}) Message {
	var room Room
	return Message{
		Type:    LeftType,
		Payload: room,
	}
}
