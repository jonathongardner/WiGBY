package camera

import (
	"fmt"
	"sync"
	"context"
	"encoding/base64"

	"gocv.io/x/gocv"
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
	// lock       sync.Mutex // need lock if we move camera frame capture to another go routine
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

func (h *Hub) Run(deviceID int, ctx context.Context) (error) {
	var wgClients sync.WaitGroup
	//-----------Camera Setup---------------
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		return fmt.Errorf("Device closed starting: %v (%v)", deviceID, err)
	}
	log.Infoln("Started camera")

	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	defer func() {
		log.Infof("Hub finishing running, closing %v clients", len(h.clients))
		for _, client := range h.clients {
			delete(h.clients, client.Identifier)
			close(client.send)
		}
		// wait for clients to all be closed
		wgClients.Wait()
	}()
	//-----------Camera Setup---------------

	log.Infoln("Camera hub listening for clients")
	for {
		select {
		case <- ctx.Done():
			log.Info("Closing hub...")
			return nil
		case client := <-h.register:
			wgClients.Add(1)
			go func() {
				defer wgClients.Done()
				client.writePump()
			}()
			h.clients[client.Identifier] = client
			log.Infof("%v clients connected", len(h.clients))
		case client := <-h.unregister:
			if _, ok := h.clients[client.Identifier]; ok {
				delete(h.clients, client.Identifier)
				close(client.send)
			}
			log.Infof("%v clients connected", len(h.clients))
		default:
			if ok := webcam.Read(&img); !ok {
				return fmt.Errorf("Device closed reading: %v (%v)", deviceID, err)
			}
			if img.Empty() {
				continue
			}

			buf, _ := gocv.IMEncode(".jpg", img)
			bytes := buf.GetBytes()
			base64String := base64.StdEncoding.EncodeToString(bytes)
			base64Bytes := []byte(base64String) // ~ Was 74029, Is 98708
			// log.Infof("Was %v, Is %v", len(bytes), len(base64Bytes))
			// if this is moved to its own process need to lock
			for _, client := range h.clients {
				// Select to skip streams which are sleeping to drop frames.
				// This might need more thought.
				select {
				case client.send <- base64Bytes:
				default:
				}
			}
			buf.Close()
		}
	}
}
