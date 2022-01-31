package main

type MessageType string

const (
	// Message types received from clients
	ConnectType MessageType = "connect"
	CreateType              = "create"
	JoinType                = "join"
	LeaveType               = "leave"

	// Message types sent from server
	RoomsType   = "rooms"
	CreatedType = "created"
	JoinedType  = "joined"
	LeftType    = "left"
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
}
