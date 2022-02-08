package main

/*
  Client represents servicees
*/
type Client struct {
	Name string `json:"name"`
	Host bool   `json:"host"`
	Room *Room  `json:"room"`
	// potentially an IP?
}
