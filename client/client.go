package client

import (
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID         int64
	Connection *websocket.Conn
	InputChan  chan *Message
	OutputChan chan *Event
	ErrorChan  chan *Message
}

func New(id int64, connection *websocket.Conn) *Client {
	return &Client{id, connection, make(chan *Message, 10), make(chan *Event, 10), make(chan *Message, 10)}
}

func (c *Client) Run() {
	for {
		select {
		case e := <-c.OutputChan:
			bytes, err := json.Marshal(e)
			if err != nil {
				log.Fatal(err)
			}

			c.Connection.WriteMessage(websocket.TextMessage, bytes)
			break
		}
	}
}
