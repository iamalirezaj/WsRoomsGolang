package main

import (
	"log"
	"github.com/gorilla/websocket"
)

/* To figure out if they wanna broadcast to all or broadcast to all except them */
type Message struct {
	msg   []byte
}

/* Reads and writes messages from client */
type Client struct {
	conn *websocket.Conn
	out  chan Message
}

/* Reads and pumps to out channel */
func (c *Client) ReadLoop() {
	defer close(c.out)
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.out <- Message{message}
	}
}

/* Writes a message to the client */
func (c *Client) WriteMessage(msg []byte) {
	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
	}
}

/* Constructor */
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:   conn,
		out:    make(chan Message),
	}
}