package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

type Hub struct {
	Rooms   map[string]*Room         `json:"rooms"`
	lobby   map[*websocket.Conn]bool `json:"-"`
	channel chan Message             `json:"-"`
	rdb     *redis.Client            `json:"-"`
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
			delete(h.lobby, conn)
			break
		}
		switch msg.Type {
		case ConnectType:
			h.channel <- h.GetRooms(ctx)
		case CreateType:
			h.channel <- h.NewRoom(ctx, msg)
			h.channel <- h.GetRooms(ctx)
		case JoinType:
			h.channel <- h.JoinRoom(ctx, msg)
		case LeaveType:
			h.channel <- h.LeaveRoom(ctx, msg.Payload)
		}
	}
}

/*
	Broadcast receives Messages from the hub's Channel and writes the message to each
	connection within the Lobby
*/
func (h *Hub) Broadcast() {
	for {
		msg := <-h.channel
		log.Println("Broadcasting: ", msg)
		for conn, ok := range h.lobby {

			// here we can have some sort of middleware that prevents
			// certain events from firing to specific clients
			if err := conn.WriteJSON(&msg); !ok || (err != nil && !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF) {
				log.Println("Unmarshal error: ", err)
				conn.Close()
				delete(h.lobby, conn)
				continue
			}
		}
	}
}

func (h *Hub) GetRooms(ctx context.Context) Message {
	var names []string
	for roomname := range h.Rooms {
		names = append(names, roomname)
	}
	return Message{
		Type:    RoomsType,
		Payload: names,
	}
}

func (h *Hub) NewRoom(ctx context.Context, message Message) Message {
	var room Room
	message.DecodePayload(&room)
	if _, ok := h.Rooms[room.Name]; ok {
		return Message{
			Type:    "error",
			Payload: "Room name taken",
		}
	}
	room.Hub = h
	room.Clients = make(map[*Client]bool)
	room.Channel = make(chan Message)
	h.rdb.HSet(ctx, "room:"+room.Name, "name", room.Name, "hostname", room.Hostname)
	h.rdb.RPush(ctx, "clients:"+room.Name, room.Hostname)
	h.Rooms[room.Name] = &room
	return Message{
		Type:    CreatedType,
		Payload: room,
	}
}

func (h *Hub) JoinRoom(ctx context.Context, msg Message) Message {
	/* if the user is the first to join the room, implicitly we will assume
	this is the host, because otherwise, this is some sort of bot that is
	listening for new rooms and hacking

	lets start by defining the payload what is the relevant information that
	we should be receiving in order to join?

	we should at least know the name of the room
	otherwise we should know everything that we need to know from the hub
	we also need to know the name of the client who is attempting to join

	while checking for the room, we'll also check the map of clients within
	the room, to assign the user as the host (if need be)

	*/
	var room Room
	var roomname string
	msg.DecodePayload(&roomname)
	if h.Rooms[roomname] == nil {
		log.Fatal("attempt to join a nil room")
	}

	return Message{
		Type:    JoinedType,
		Payload: room,
	}
}

func (h *Hub) LeaveRoom(ctx context.Context, payload interface{}) Message {
	var room Room
	log.Fatal("Leave Room is not a service developed")
	return Message{
		Type:    LeftType,
		Payload: room,
	}
}
