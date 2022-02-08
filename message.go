package main

import (
	"encoding/json"
	"log"
)

type MessageType string

const (
	// Message types received from clients
	ConnectType MessageType = "connect"
	CreateType              = "create"
	JoinType                = "join"
	LeaveType               = "leave"

	// Message types sent from server
	RoomsType     = "rooms"
	CreatedType   = "created"
	JoinedType    = "joined"
	LeftType      = "left"
	ErrorType     = "error"
	ConnectedType = "connected"
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
}

func (m *Message) DecodePayload(dst interface{}) {
	bytes, err := json.Marshal(m.Payload)
	if err != nil {
		log.Println("Could not decode payload, marshal err: ", err)
	}
	if err := json.Unmarshal(bytes, dst); err != nil {
		log.Println("Could not decode payload, unmarshal err: ", err)
	}
}
