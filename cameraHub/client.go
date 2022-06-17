package cameraHub

import (
	"time"

	"github.com/gorilla/websocket"
	// log "github.com/sirupsen/logrus"
)


const (
	// Time allowed to write a message to the peer.
	writeWait = 2 * time.Second

	// time between sending frames
	frameInterval = 50 * time.Millisecond
	// frameInterval = 5 * time.Second
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// channel of outbound frame.
	send chan []byte

	// Name of this user (can be empty)
	// IPAddress of client
	// Identifier: Unique
	Name       string `json:"name"`
	IPAddress  string `json:"ipAddress"`
	Identifier string `json:"identifier"`
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	defer c.conn.Close()
	for {
		// only send fram every x Milliseconds
		time.Sleep(frameInterval)

		img, ok := <-c.send
		c.conn.SetWriteDeadline(time.Now().Add(writeWait)) // 2 seconds

		if !ok {
			// The hub closed the channel. So we dont need to delete
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			c.delete()
			return
		}
		w.Write(img)

		if err := w.Close(); err != nil {
			c.delete()
			return
		}
	}
}

func (c *Client) delete() {
	c.hub.unregister <- c
}
