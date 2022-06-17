package cameraHub

import (
  "sync"

  "github.com/google/uuid"
  "github.com/gorilla/websocket"
  log "github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	clients    map[string]*Client // Registered clients.
	register   chan *Client // Register requests from the clients.
	unregister chan *Client // Unregister requests from clients.
  lock       sync.Mutex // need lock to iterate over map?
}

func NewHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) NewClient(conn *websocket.Conn, ipAddress string, name string) (*Client) {
  id := uuid.New().String()
  if name == "" {
    name = id
  }
  cl := &Client{hub: h, conn: conn, send: make(chan []byte), Name: name, IPAddress: ipAddress, Identifier: id}

  // use channel so thread safe?
  h.register <- cl

  return cl
}

func (h *Hub) Run() {
  log.Infoln("Camera hub listening for clients")
	for {
		select {
		case client := <-h.register:
      h.lock.Lock()
			h.clients[client.Identifier] = client
      h.lock.Unlock()
		case client := <-h.unregister:
			if _, ok := h.clients[client.Identifier]; ok {
        h.lock.Lock()
				delete(h.clients, client.Identifier)
        h.lock.Unlock()
				close(client.send)
			}
      // TODO listen for errors here and close everything if error
		}
    log.Infof("%v clients connected", len(h.clients))
	}
}
